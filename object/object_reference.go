package object

type Reference interface {
	Type() Type
	Inspect() string
	UpdateValue(Object) (Object, error)
	GetValue() Object
}

type VariableReference struct {
	Env  *Environment
	Name string
	ReferenceType
}

type ReferenceType struct {
	IsConstantReference bool
	ValueType           Type
}

type ArrayElementReference struct {
	*Array
	Index int64
	ReferenceType
}

type HashElementReference struct {
	*Hash
	Key Object
	ReferenceType
}

func (vr *VariableReference) UpdateValue(value Object) (Object, error) {
	return vr.Env.Set(vr.Name, value)
}

func (vr *VariableReference) GetValue() Object {
	return vr.Env.Get(vr.Name)
}

func (vr *VariableReference) Type() Type {
	return ReferenceType{vr.ReferenceType.IsConstantReference, vr.GetValue().Type()}
}

func (vr *VariableReference) Inspect() string {
	return vr.GetValue().Inspect()
}

func (ar *ArrayElementReference) UpdateValue(value Object) (Object, error) {
	ar.Array.Elements[ar.Index] = value
	return value, nil
}

func (ar *ArrayElementReference) GetValue() Object {
	return ar.Array.Elements[ar.Index]
}

func (ar *ArrayElementReference) Type() Type {
	return ReferenceType{ar.ReferenceType.IsConstantReference, ar.GetValue().Type()}
}

func (ar *ArrayElementReference) Inspect() string {
	return ar.GetValue().Inspect()
}

func (hr *HashElementReference) UpdateValue(value Object) (Object, error) {
	hashKey := hr.Key.(Hashable).HashKey()
	hr.Hash.Pairs[hashKey] = HashPair{hr.Key, value}
	return value, nil
}

func (hr *HashElementReference) GetValue() Object {
	hashKey := hr.Key.(Hashable).HashKey()
	return hr.Hash.Pairs[hashKey].Value
}

func (hr *HashElementReference) Type() Type {
	return ReferenceType{hr.ReferenceType.IsConstantReference, hr.GetValue().Type()}
}

func (hr *HashElementReference) Inspect() string {
	return hr.GetValue().Inspect()
}

func (t ReferenceType) Signature() string             { return "&" + t.ValueType.Signature() }
func (t ReferenceType) Kind() Kind                    { return t.ValueType.Kind() }
func (t ReferenceType) Builtins() *FunctionRepository { return nil }
func (t ReferenceType) IsConstant() bool              { return t.IsConstantReference }
