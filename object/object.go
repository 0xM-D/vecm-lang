package object

type Type interface {
	Signature() string
	Kind() Kind
	Builtins() *FunctionRepository
	IsConstant() bool
}

type Object interface {
	Type() Type
	Inspect() string
}
