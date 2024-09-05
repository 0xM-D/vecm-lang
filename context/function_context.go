package context

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/util"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type FunctionContext struct {
	sharedContextProperties SharedContextProperties
	functionParams          *VariableStore
	*ir.Func
}

func (ctx *FunctionContext) GetParentContext() Context {
	return ctx.sharedContextProperties.parentContext
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

func (ctx *FunctionContext) LookUpIdentifier(name string) (Variable, bool) {
	variable, ok := ctx.functionParams.LookUpVariable(name)
	if ok {
		return variable, ok
	}

	if ctx.GetParentContext() == nil {
		return nil, false
	}

	return ctx.GetParentContext().LookUpIdentifier(name)
}

func (ctx *FunctionContext) GetFunction(name string, params ...*ir.Param) *ir.Func {
	return ctx.GetParentContext().GetFunction(name, params...)
}

func (ctx *FunctionContext) DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func {
	return ctx.GetParentContext().DeclareFunction(name, retType, params...)
}

func (ctx *FunctionContext) DeclareLocalVariable(_ string, _ types.Type) *ir.InstAlloca {
	return nil // TODO: Throw error
}
