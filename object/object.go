package object

type ObjectType interface {
	Signature() string
}

type Object interface {
	Type() ObjectType
	Inspect() string
}
