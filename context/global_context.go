package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type GlobalContext struct {
	Module *ir.Module
	sharedContextProperties SharedContextProperties
}

func NewGlobalContext(parentContext *Context) *GlobalContext {
	return &GlobalContext{
		Module: ir.NewModule(),
		// sharedContextProperties: SharedContextProperties{
		// 	functionStore: FunctionStore{map[string]*FunctionObject{}},
		// },
	}
}

func (ctx *GlobalContext) GetParentContext() *Context {
	return ctx.sharedContextProperties.parentContext;
}

func (ctx *GlobalContext) GetFunction(signature types.FuncType) (*ir.Func, bool) {
	funcs := ctx.Module.Funcs
	for _, f := range(funcs) {
		if f.Sig.Equal(&signature) {
			return f, true
		}
	}
	return nil, false;
}

func (ctx GlobalContext) DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func {
	fn := ctx.Module.NewFunc(name, retType, params...)
	return fn;
}

