package object

type ObjectKind string

var intrinsicTypeFunctionRepositories = map[ObjectKind]*FunctionRepository{}

func (o ObjectKind) Signature() string             { return string(o) }
func (o ObjectKind) Kind() ObjectKind              { return o }
func (o ObjectKind) Builtins() *FunctionRepository { return intrinsicTypeFunctionRepositories[o] }

const (
	Invalid      ObjectKind = "invalid"
	IntegerKind  ObjectKind = "int"
	BooleanKind  ObjectKind = "bool"
	StringKind   ObjectKind = "string"
	ArrayKind    ObjectKind = "array"
	HashKind     ObjectKind = "hash"
	FunctionKind ObjectKind = "function"
	NullKind     ObjectKind = "null"
	VoidKind     ObjectKind = "void"
	ErrorKind    ObjectKind = "error"
	AnyKind      ObjectKind = "any"
)
