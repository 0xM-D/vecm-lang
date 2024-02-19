package context

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

func NewGlobalContext() *GlobalContext {
	return &GlobalContext{
		Module: ir.NewModule(),
		BlockContext: BlockContext{
			FunctionStore: FunctionStore{map[string]FunctionObject{}},
		},
	}
}

func (ctx *GlobalContext) DeclareFunction(name string, fn *ir.Func) error {
	if _, exists := ctx.Fns[name]; exists {
		return fmt.Errorf("function %s already exists", name);
	}
	ctx.Fns[name] = FunctionObject{Name: name, Fn: fn};

	return nil;
}