package object

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
	AnyKind             ObjectKind = "any"
)

func initIntrinsicTypeBuiltins() map[ObjectKind]*FunctionRepository {
	repos := map[ObjectKind]*FunctionRepository{}

	numberBuiltins := initNumberBuiltins()
	for _, nt := range NumberTypes {
		repos[nt] = numberBuiltins
	}
	repos[ArrayKind] = initArrayBuiltins()
	repos[StringKind] = initStringBuiltins()

	return repos
}

func initNumberBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("toString", FunctionObjectType{ParameterTypes: []FunctionParameterType{}, ReturnValueType: StringKind}, numberToString)

	return &repo
}

func initStringBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("length", FunctionObjectType{ParameterTypes: []FunctionParameterType{}, ReturnValueType: Int64Kind}, stringLength)

	return &repo
}

func initArrayBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("size", FunctionObjectType{ParameterTypes: []FunctionParameterType{}, ReturnValueType: Int64Kind}, arraySize)
	repo.register("push", FunctionObjectType{ParameterTypes: []FunctionParameterType{{AnyKind, false}}, ReturnValueType: ArrayKind}, arrayPush)
	repo.register("pushMultiple", FunctionObjectType{ParameterTypes: []FunctionParameterType{{AnyKind, false}, {Int64Kind, false}}, ReturnValueType: ArrayKind}, arrayPushMultiple)
	repo.register("delete", FunctionObjectType{ParameterTypes: []FunctionParameterType{{Int64Kind, false}, {Int64Kind, false}}, ReturnValueType: ArrayKind}, arrayDelete)
	repo.register("slice", FunctionObjectType{ParameterTypes: []FunctionParameterType{{Int64Kind, false}, {Int64Kind, false}}, ReturnValueType: ArrayKind}, arraySlice)
	return &repo
}

func numberToString(params ...Object) Object {
	number := params[0].(*Number)
	return &String{number.Inspect()}
}

func stringLength(params ...Object) Object {
	str := params[0].(*String)
	return &Number{Value: uint64(len(str.Value)), Kind: UInt64Kind}
}

func arraySize(params ...Object) Object {
	arr := params[0].(*Array)
	return &Number{Value: uint64(len(arr.Elements)), Kind: UInt64Kind}
}

func arrayPush(params ...Object) Object {
	arr := params[0].(*Array)
	elem := params[1]

	arr.Elements = append(arr.Elements, elem)
	return arr
}

func arrayDelete(params ...Object) Object {
	arr := params[0].(*Array)
	startIndex := UnwrapReferenceObject(params[1]).(*Number).GetInt64()
	count := UnwrapReferenceObject(params[2]).(*Number).GetInt64()
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
	size := UnwrapReferenceObject(params[2]).(*Number).GetInt64()

	newElements := make([]Object, 0, size)
	for i := int64(0); i != size; i++ {
		newElements = append(newElements, element)
	}

	arr.Elements = append(arr.Elements, newElements...)

	return arr
}

func arraySlice(params ...Object) Object {
	arr := params[0].(*Array)
	startIndex := UnwrapReferenceObject(params[1]).(*Number).GetInt64()
	count := UnwrapReferenceObject(params[2]).(*Number).GetInt64()

	startIndex = boundArrayIndex(arr, startIndex)
	endIndex := boundArrayIndex(arr, startIndex+count-1)

	return &Array{ArrayObjectType: arr.ArrayObjectType, Elements: arr.Elements[startIndex : endIndex+1]}
}

func boundArrayIndex(arr *Array, index int64) int64 {
	if index < 0 {
		index = 0
	}

	if index >= int64(len(arr.Elements)) {
		index = int64(len(arr.Elements)) - 1
	}

	return index

}
