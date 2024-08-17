package context

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type VariableStore struct {
	variables map[string]Variable
}

type Variable interface {
	GetName() string
	GetType() types.Type
	GetValue() value.Value
}

type VariableImpl struct {
	name string
	t types.Type
	value value.Value
}

func (v *VariableImpl) GetName() string {
	return v.name
}

func (v *VariableImpl) GetType() types.Type {
	return v.t
}

func (v *VariableImpl) GetValue() value.Value {
	return v.value
}

func (vs *VariableStore) DeclareVariable(name string, t types.Type, value value.Value) Variable{
	vs.variables[name] = &VariableImpl{name, t, value}
	return vs.variables[name]
}

func (vs *VariableStore) LookUpVariable(name string) (Variable, bool) {
	variable, ok := vs.variables[name]
	return variable, ok
}

func NewVariableStore() *VariableStore {
	return &VariableStore{make(map[string]Variable)}
}