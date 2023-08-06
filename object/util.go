package object

func IsInteger(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*IntegerObjectType)
	return ok
}

func IsBoolean(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*BooleanObjectType)
	return ok
}

func IsNull(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*NullObjectType)
	return ok
}

func IsString(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*StringObjectType)
	return ok
}

func IsError(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*ErrorObjectType)
	return ok
}

func IsArray(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*ArrayObjectType)
	return ok
}

func IsHash(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(*HashObjectType)
	return ok
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
