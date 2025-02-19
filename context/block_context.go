package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

type BlockContext struct {
	sharedContextProperties SharedContextProperties
	variableStore           *VariableStore
	*ir.Block
}

func (ctx *BlockContext) GetParentContext() (Context, error) {
	if ctx.sharedContextProperties.parentContext == nil {
		return nil, errors.New("reached the top of the context chain")
	}
	return ctx.sharedContextProperties.parentContext, nil
}

func (ctx *BlockContext) GetParentFunctionContext() (*FunctionContext, error) {
	var currentContext Context = ctx
	var err error

	// go up the chain until we find a function context
	for currentContext != nil {
		if functionContext, ok := currentContext.(*FunctionContext); ok {
			return functionContext, nil
		}
		currentContext, err = currentContext.GetParentContext()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get parent function context")
		}
	}

	return nil, errors.New("failed to get parent function context: no function context found")
}

func (ctx *BlockContext) LookUpIdentifier(name string) (Variable, error) {
	variable, ok := ctx.variableStore.LookUpVariable(name)
	if ok {
		return variable, nil
	}

	parentContext, err := ctx.GetParentContext()
	if err != nil {
		return nil, errors.Wrap(err, "failed to look up identifier")
	}

	identifier, err := parentContext.LookUpIdentifier(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to look up identifier")
	}

	return identifier, nil
}

func (ctx *BlockContext) GetFunction(name string, params ...*ir.Param) (*ir.Func, error) {
	parentContext, err := ctx.GetParentContext()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get function")
	}

	function, err := parentContext.GetFunction(name, params...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get function")
	}

	return function, nil
}

func (ctx *BlockContext) DeclareFunction(
	name string,
	retType types.Type,
	isVariadic bool,
	params ...*ir.Param,
) (*ir.Func, error) {
	parentContext, err := ctx.GetParentContext()
	if err != nil {
		return nil, errors.Wrap(err, "failed to declare function")
	}

	function, err := parentContext.DeclareFunction(name, retType, isVariadic, params...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to declare function")
	}

	return function, nil
}

func (ctx *BlockContext) DeclareLocalVariable(name string, t types.Type) (*ir.InstAlloca, error) {
	alloca := ctx.Block.NewAlloca(t)
	ctx.variableStore.DeclareLocalVariable(name, alloca)
	return alloca, nil
}

func NewBlockContext(parentContext Context, block *ir.Block) *BlockContext {
	return &BlockContext{
		sharedContextProperties: SharedContextProperties{parentContext},
		variableStore:           NewVariableStore(),
		Block:                   block,
	}
}
