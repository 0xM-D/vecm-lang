package object

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

func FUNCTION_OBJ() FunctionObjectType {
	return FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: INTEGER_OBJ()}
}

func RETURN_VALUE_OBJ() ReturnValueObjectType {
	return ReturnValueObjectType{ReturnType: INTEGER_OBJ()}
}
