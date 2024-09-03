package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type VariableStore struct {
	variables map[string]Variable
}

type Variable interface {
	GetName() string
	GetType() types.Type
	GetAddress() value.Value
}

type FunctionParamVariable struct {
	name  string
	param *ir.Param
}

type LocalVariable struct {
	name   string
	alloca *ir.InstAlloca
}

func (v *LocalVariable) GetName() string {
	return v.name
}

func (v *LocalVariable) GetType() types.Type {
	return v.alloca.ElemType
}

func (v *LocalVariable) GetAddress() value.Value {
	return v.alloca
}

func (v *FunctionParamVariable) GetName() string {
	return v.name
}

func (v *FunctionParamVariable) GetType() types.Type {
	return v.param.Type()
}

func (v *FunctionParamVariable) GetAddress() value.Value {
	return v.param
}

func (vs *VariableStore) DeclareLocalVariable(name string, alloca *ir.InstAlloca) Variable {
	vs.variables[name] = &LocalVariable{name, alloca}
	return vs.variables[name]
}

func (vs *VariableStore) DeclareFunctionParam(name string, param *ir.Param) Variable {
	vs.variables[name] = &FunctionParamVariable{name, param}
	return vs.variables[name]
}

func (vs *VariableStore) LookUpVariable(name string) (Variable, bool) {
	variable, ok := vs.variables[name]
	return variable, ok
}

func NewVariableStore() *VariableStore {
	return &VariableStore{make(map[string]Variable)}
}
