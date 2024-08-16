package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Context interface {
	GetParentContext() Context
	GetFunction(signature types.FuncType) (*ir.Func, bool)
	DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func
	DeclareLocalVariable(name string, t types.Type) *ir.InstAlloca
	LookUpIdentifier(name string) (Variable, bool)
}

type SharedContextProperties struct {	
	parentContext Context
	// functionStore FunctionStore
}
