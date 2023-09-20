package object

import (
	"strconv"
)

type ObjectKind string

var intrinsicTypeFunctionRepositories = initIntrinsicTypeBuiltins()

func (o ObjectKind) Signature() string             { return string(o) }
func (o ObjectKind) Kind() ObjectKind              { return o }
func (o ObjectKind) Builtins() *FunctionRepository { return intrinsicTypeFunctionRepositories[o] }
func (o ObjectKind) IsConstant() bool              { return true }

const (
	Invalid             ObjectKind = "invalid"
	Int8Kind            ObjectKind = "int8"
	Int16Kind           ObjectKind = "int16"
	Int32Kind           ObjectKind = "int32"
	Int64Kind           ObjectKind = "int64"
	UInt8Kind           ObjectKind = "uint8"
	UInt16Kind          ObjectKind = "uint16"
	UInt32Kind          ObjectKind = "uint32"
	UInt64Kind          ObjectKind = "uint64"
	Float32Kind         ObjectKind = "float32"
	Float64Kind         ObjectKind = "float64"
	BooleanKind         ObjectKind = "bool"
	StringKind          ObjectKind = "string"
	ArrayKind           ObjectKind = "array"
	HashKind            ObjectKind = "hash"
	FunctionKind        ObjectKind = "function"
	BuiltinFunctionKind ObjectKind = "builtinFunction"
	NullKind            ObjectKind = "null"
	VoidKind            ObjectKind = "void"
	ErrorKind           ObjectKind = "error"
	AnyKind             ObjectKind = "any"
)

func initIntrinsicTypeBuiltins() map[ObjectKind]*FunctionRepository {
	repos := map[ObjectKind]*FunctionRepository{}

	repos[Int64Kind] = initIntegerBuiltins()
	repos[ArrayKind] = initArrayBuiltins()
	repos[StringKind] = initStringBuiltins()

	return repos
}

func initIntegerBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("toString", FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: StringKind}, intToString)

	return &repo
}

func initStringBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("length", FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: Int64Kind}, stringLength)

	return &repo
}

func initArrayBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("size", FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: Int64Kind}, arraySize)
	repo.register("push", FunctionObjectType{ParameterTypes: []ObjectType{Int64Kind}, ReturnValueType: ArrayKind}, arrayPush)
	repo.register("pushMultiple", FunctionObjectType{ParameterTypes: []ObjectType{AnyKind, Int64Kind}, ReturnValueType: ArrayKind}, arrayPushMultiple)
	repo.register("delete", FunctionObjectType{ParameterTypes: []ObjectType{Int64Kind, Int64Kind}, ReturnValueType: ArrayKind}, arrayDelete)
	repo.register("slice", FunctionObjectType{ParameterTypes: []ObjectType{Int64Kind, Int64Kind}, ReturnValueType: ArrayKind}, arraySlice)
	return &repo
}

func intToString(params ...Object) Object {
	integer := params[0].(*Number[int64])
	return &String{strconv.FormatInt(integer.Value, 10)}
}

func stringLength(params ...Object) Object {
	str := params[0].(*String)
	return &Number[int64]{int64(len(str.Value))}
}

func arraySize(params ...Object) Object {
	arr := params[0].(*Array)
	return &Number[int64]{int64(len(arr.Elements))}
}

func arrayPush(params ...Object) Object {
	arr := params[0].(*Array)
	elem := params[1]

	arr.Elements = append(arr.Elements, elem)
	return arr
}

func arrayDelete(params ...Object) Object {
	arr := params[0].(*Array)
	startIndex := UnwrapReferenceObject(params[1]).(*Number[int64]).Value
	count := UnwrapReferenceObject(params[2]).(*Number[int64]).Value
	arrLen := int64(len(arr.Elements))

	if startIndex+count >= arrLen {
		if startIndex >= arrLen {
			arr.Elements = []Object{}
		} else {
			arr.Elements = arr.Elements[:startIndex]
		}
	} else {
		arr.Elements = append(arr.Elements[:startIndex], arr.Elements[startIndex+count:]...)
	}

	return arr
}

func arrayPushMultiple(params ...Object) Object {
	arr := params[0].(*Array)
	element := UnwrapReferenceObject(params[1])
	size := int(UnwrapReferenceObject(params[2]).(*Number[int64]).Value)

	newElements := make([]Object, 0, size)
	for i := 0; i != size; i++ {
		newElements = append(newElements, element)
	}

	arr.Elements = append(arr.Elements, newElements...)

	return arr
}

func arraySlice(params ...Object) Object {
	arr := params[0].(*Array)
	startIndex := int(UnwrapReferenceObject(params[1]).(*Number[int64]).Value)
	count := int(UnwrapReferenceObject(params[2]).(*Number[int64]).Value)

	startIndex = boundArrayIndex(arr, startIndex)
	endIndex := boundArrayIndex(arr, startIndex+count-1)

	return &Array{ArrayObjectType: arr.ArrayObjectType, Elements: arr.Elements[startIndex : endIndex+1]}
}

func boundArrayIndex(arr *Array, index int) int {
	if index < 0 {
		index = 0
	}

	if index >= len(arr.Elements) {
		index = len(arr.Elements) - 1
	}

	return index

}
