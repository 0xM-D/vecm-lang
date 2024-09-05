package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type GlobalContext struct {
	Module        *ir.Module
	variableStore *VariableStore
}

func NewGlobalContext() *GlobalContext {
	return &GlobalContext{
		Module:        ir.NewModule(),
		variableStore: NewVariableStore(),
	}
}

func (ctx *GlobalContext) GetParentContext() Context {
	return nil
}

func (ctx GlobalContext) GetFunction(name string, params ...*ir.Param) *ir.Func {
	funcs := ctx.Module.Funcs
	for _, f := range funcs {
		if f.Name() == name {
			return f
		}
	}
	return nil
}

func (ctx GlobalContext) DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func {
	fn := ctx.Module.NewFunc(name, retType, params...)
	return fn
}

func (ctx *GlobalContext) DeclareLocalVariable(name string, t types.Type) *ir.InstAlloca {
	return nil // TODO: Throw error
}

func (ctx *GlobalContext) LookUpIdentifier(name string) (Variable, bool) {
	variable, ok := ctx.variableStore.LookUpVariable(name)
	if ok {
		return variable, ok
	}

	return nil, false
}
