package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Context interface {
	GetParentContext() (Context, error)
	GetFunction(name string, params ...*ir.Param) (*ir.Func, error)
	DeclareFunction(name string, retType types.Type, isVariadic bool, params ...*ir.Param) (*ir.Func, error)
	DeclareLocalVariable(name string, t types.Type) (*ir.InstAlloca, error)
	LookUpIdentifier(name string) (Variable, error)
}

type SharedContextProperties struct {
	parentContext Context
	// functionStore FunctionStore
}
