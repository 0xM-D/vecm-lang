package object

var INTEGER_TYPES = map[ObjectKind]bool{
	Int8Kind:   true,
	Int16Kind:  true,
	Int32Kind:  true,
	Int64Kind:  true,
	UInt8Kind:  true,
	UInt16Kind: true,
	UInt32Kind: true,
	UInt64Kind: true,
}

func IsInteger(i Object) bool {
	return IsIntegerKind(i.Type().Kind())
}

func IsFloat32(i Object) bool {
	return i.Type().Kind() == Float32Kind
}

func IsFloat64(i Object) bool {
	return i.Type().Kind() == Float64Kind
}

func IsFloat(i Object) bool {
	return IsFloat32(i) || IsFloat64(i)
}

func IsNumber(i Object) bool {
	return IsInteger(i) || IsFloat(i)
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

func IsBuiltinFunction(i Object) bool {
	return i.Type().Kind() == BuiltinFunctionKind
}

func IsReturnValue(i Object) bool {
	_, ok := UnwrapReferenceObject(i).Type().(ReturnValueObjectType)
	return ok
}

func TypesMatch(t1 ObjectType, t2 ObjectType) bool {
	return t1.Signature() == t2.Signature()
}

func IsIntegerKind(k ObjectKind) bool {
	_, isInteger := INTEGER_TYPES[k]
	return isInteger
}

func IsNumberKind(k ObjectKind) bool {
	return k == Float32Kind || k == Float64Kind || IsIntegerKind(k)
}

func UnwrapReferenceObject(or Object) Object {
	object, ok := or.(ObjectReference)
	if ok {
		return object.GetValue()
	}
	return or
}

func UnwrapReferenceType(t ObjectType) ObjectType {
	reference, ok := t.(ReferenceType)
	if ok {
		return reference.ValueType
	}
	return t
}
