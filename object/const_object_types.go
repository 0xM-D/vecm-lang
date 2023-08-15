package object

type AnyObjectType struct{}

func (*AnyObjectType) Signature() string { return "any" }

func STRING_OBJ() ObjectType {
	return &StringObjectType{}
}

func BOOLEAN_OBJ() ObjectType {
	return &BooleanObjectType{}
}

func NULL_OBJ() ObjectType {
	return &NullObjectType{}
}

func INTEGER_OBJ() ObjectType {
	return &IntegerObjectType{}
}

func ERROR_OBJ() ObjectType {
	return &ErrorObjectType{}
}

func HASH_OBJ() HashObjectType {
	return HashObjectType{KeyType: STRING_OBJ(), ValueType: STRING_OBJ()}
}

func ARRAY_OBJ() ArrayObjectType {
	return ArrayObjectType{ElementType: INTEGER_OBJ()}
}

func VOID_OBJ() VoidObjectType {
	return VoidObjectType{}
}
