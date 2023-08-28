package object

type ObjectType interface {
	Signature() string
	Kind() ObjectKind
	Builtins() *FunctionRepository
	IsConstant() bool
}

type Object interface {
	Type() ObjectType
	Inspect() string
}
