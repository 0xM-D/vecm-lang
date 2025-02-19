package context

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
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

func (ctx *GlobalContext) GetParentContext() (Context, error) {
	return nil, errors.New("global context has no parent context")
}

func (ctx GlobalContext) GetFunction(name string, _ ...*ir.Param) (*ir.Func, error) {
	funcs := ctx.Module.Funcs
	for _, f := range funcs {
		if f.Name() == name {
			return f, nil
		}
	}
	return nil, fmt.Errorf("function %s not found", name)
}

func (ctx GlobalContext) DeclareFunction(
	name string,
	retType types.Type,
	isVariadic bool,
	params ...*ir.Param,
) (*ir.Func, error) {
	fn := ctx.Module.NewFunc(name, retType, params...)
	fn.Sig.Variadic = isVariadic
	return fn, nil
}

func (ctx *GlobalContext) DeclareLocalVariable(_ string, _ types.Type) (*ir.InstAlloca, error) {
	return nil, errors.New("cannot declare local variable in global context")
}

func (ctx *GlobalContext) LookUpIdentifier(name string) (Variable, error) {
	variable, ok := ctx.variableStore.LookUpVariable(name)
	if ok {
		return variable, nil
	}

	return nil, fmt.Errorf("identifier %s not found", name)
}
