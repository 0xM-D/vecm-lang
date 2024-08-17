package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type BlockContext struct {
	sharedContextProperties SharedContextProperties
	variableStore *VariableStore
	*ir.Block
}

func (ctx *BlockContext) GetParentContext() Context {
	return ctx.sharedContextProperties.parentContext;
}

func (ct *BlockContext) GetParentFunctionContext() *FunctionContext {
	
	// go up the chain until we find a function context
	for ctx := ct.GetParentContext(); ctx != nil; ctx = ctx.GetParentContext() {
		if functionContext, ok := ctx.(*FunctionContext); ok {
			return functionContext
		}
	}

	// This should never happen, TODO: Throw compiler error
	return nil
}

func (ctx *BlockContext) LookUpIdentifier(name string) (Variable, bool) {
	variable, ok := ctx.variableStore.LookUpVariable(name)
	if ok {
		return variable, ok
	}

	if ctx.GetParentContext() == nil {
		return nil, false
	}

	return ctx.GetParentContext().LookUpIdentifier(name)
}

func (ctx *BlockContext) GetFunction(signature types.FuncType) (*ir.Func, bool) {
	return ctx.GetParentContext().GetFunction(signature)
}

func (ctx *BlockContext) DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func {
	return ctx.GetParentContext().DeclareFunction(name, retType, params...)
}

func (ctx *BlockContext) DeclareLocalVariable(name string, t types.Type) *ir.InstAlloca {

	alloca := ctx.Block.NewAlloca(t)
	ctx.variableStore.DeclareVariable(name, t, alloca)
	return alloca
}

func NewBlockContext(parentContext Context, block *ir.Block) *BlockContext {
	return &BlockContext{
		sharedContextProperties: SharedContextProperties{parentContext},
		variableStore: NewVariableStore(),
		Block: block,
	}
}