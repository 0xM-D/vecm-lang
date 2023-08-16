package object

func IsInteger(i Object) bool {
	return i.Type().Kind() == IntegerKind
}

func IsBoolean(i Object) bool {
	return i.Type().Kind() == BooleanKind

}

func IsNull(i Object) bool {
	return i.Type().Kind() == NullKind
}

func IsString(i Object) bool {
	return i.Type().Kind() == StringKind
}

func IsError(i Object) bool {
	return i.Type().Kind() == ErrorKind
}

func IsArray(i Object) bool {
	return i.Type().Kind() == ArrayKind
}

func IsHash(i Object) bool {
	return i.Type().Kind() == HashKind
}

func IsFunction(i Object) bool {
	return i.Type().Kind() == FunctionKind
}

func IsReturnValue(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(ReturnValueObjectType)
	return ok
}

func TypesMatch(t1 ObjectType, t2 ObjectType) bool {
	return t1.Signature() == t2.Signature()
}

func UnwrapReferenceObject(or Object) Object {
	object, ok := or.(*ObjectReference)
	if ok {
		return object.Object
	}
	return or
}
