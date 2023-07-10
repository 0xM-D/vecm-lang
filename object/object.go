package object

type ObjectType string
type ObjectValue interface {
	Type() ObjectType
	Inspect() string
}
type Object struct {
	Value      ObjectValue
	IsConstant bool
}

const (
	INTEGER_OBJ      = "int"
	BOOLEAN_OBJ      = "bool"
	NULL_OBJ         = "null"
	RETURN_VALUE_OBJ = "returnvalue"
	ERROR_OBJ        = "error"
	FUNCTION_OBJ     = "function"
	STRING_OBJ       = "string"
	ARRAY_OBJ        = "array"
	HASH_OBJ         = "hash"
)
