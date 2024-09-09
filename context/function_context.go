package context

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/util"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

type FunctionContext struct {
	sharedContextProperties SharedContextProperties
	functionParams          *VariableStore
	*ir.Func
}

func (ctx *FunctionContext) GetParentContext() (Context, error) {
	if ctx.sharedContextProperties.parentContext == nil {
		return nil, errors.New("Reached the top of the context chain")
	}
	return ctx.sharedContextProperties.parentContext, nil
}

func NewFunctionContext(
	parent Context,
	fn *ir.Func,
	parameterNames []*ast.Identifier,
	parameterTypes []ast.Type,
) *FunctionContext {
	ctx := &FunctionContext{
		sharedContextProperties: SharedContextProperties{parentContext: parent},
		Func:                    fn,
		functionParams:          NewVariableStore(),
	}

	for i, name := range parameterNames {
		t, llvmTypeError := util.GetLLVMType(parameterTypes[i])
		if llvmTypeError != nil {
			panic(llvmTypeError)
		}

		if !fn.Params[i].Typ.Equal(t) {
			panic("Type mismatch")
		}

		ctx.functionParams.DeclareFunctionParam(name.Value, fn.Params[i])
	}

	return ctx
}

func (ctx *FunctionContext) LookUpIdentifier(name string) (Variable, error) {
	variable, ok := ctx.functionParams.LookUpVariable(name)
	if ok {
		return variable, nil
	}

	parentContext, err := ctx.GetParentContext()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to look up identifier")
	}

	identifier, err := parentContext.LookUpIdentifier(name)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to look up identifier")
	}

	return identifier, nil
}

func (ctx *FunctionContext) GetFunction(name string, params ...*ir.Param) (*ir.Func, error) {
	parentContext, err := ctx.GetParentContext()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get function")
	}

	function, err := parentContext.GetFunction(name, params...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get function")
	}

	return function, nil
}

func (ctx *FunctionContext) DeclareFunction(name string, retType types.Type, params ...*ir.Param) (*ir.Func, error) {
	parentContext, err := ctx.GetParentContext()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare function")
	}

	function, err := parentContext.DeclareFunction(name, retType, params...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to declare function")
	}

	return function, nil
}

func (ctx *FunctionContext) DeclareLocalVariable(_ string, _ types.Type) (*ir.InstAlloca, error) {
	return nil, errors.New("Cannot declare local variable in function context")
}
