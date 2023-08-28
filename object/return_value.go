package object

type ReturnValueObjectType struct {
	ReturnType ObjectType
}

func (r ReturnValueObjectType) Signature() string             { return r.ReturnType.Signature() }
func (r ReturnValueObjectType) Kind() ObjectKind              { return r.ReturnType.Kind() }
func (r ReturnValueObjectType) Builtins() *FunctionRepository { return r.ReturnType.Kind().Builtins() }
func (r ReturnValueObjectType) IsConstant() bool              { return true }

type ReturnValue struct {
	ReturnValueObjectType
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return rv.ReturnValueObjectType }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
