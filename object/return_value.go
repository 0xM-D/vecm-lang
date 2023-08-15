package object

type ReturnValueObjectType struct {
	ReturnType ObjectType
}

func (r ReturnValueObjectType) Signature() string { return "R" + r.ReturnType.Signature() }

type ReturnValue struct {
	ReturnValueObjectType
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return rv.ReturnValueObjectType }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
