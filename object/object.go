package object

type ObjectType interface {
	Signature() string
	Kind() ObjectKind
	Builtins() *FunctionRepository
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type ObjectReference struct {
	Object     Object
	IsConstant bool
	Identifier string
}

func (or *ObjectReference) Type() ObjectType { return or.Object.Type() }
func (or *ObjectReference) Inspect() string  { return or.Object.Inspect() }
