package object

type Kind string

var intrinsicTypeFunctionRepositories = initIntrinsicTypeBuiltins()

func (o Kind) Signature() string             { return string(o) }
func (o Kind) Kind() Kind                    { return o }
func (o Kind) Builtins() *FunctionRepository { return intrinsicTypeFunctionRepositories[o] }
func (o Kind) IsConstant() bool              { return true }

const (
	Invalid             Kind = "invalid"
	Int8Kind            Kind = "int8"
	Int16Kind           Kind = "int16"
	Int32Kind           Kind = "int32"
	Int64Kind           Kind = "int64"
	UInt8Kind           Kind = "uint8"
	UInt16Kind          Kind = "uint16"
	UInt32Kind          Kind = "uint32"
	UInt64Kind          Kind = "uint64"
	Float32Kind         Kind = "float32"
	Float64Kind         Kind = "float64"
	BooleanKind         Kind = "bool"
	StringKind          Kind = "string"
	ArrayKind           Kind = "array"
	HashKind            Kind = "hash"
	FunctionKind        Kind = "function"
	BuiltinFunctionKind Kind = "builtinFunction"
	NullKind            Kind = "null"
	VoidKind            Kind = "void"
	AnyKind             Kind = "any"
)

func initIntrinsicTypeBuiltins() map[Kind]*FunctionRepository {
	repos := map[Kind]*FunctionRepository{}

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

	repo.register(
		"toString",
		FunctionObjectType{
			ParameterTypes:  []Type{},
			ReturnValueType: StringKind,
		},
		numberToString,
	)

	return &repo
}

func initStringBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("length", FunctionObjectType{ParameterTypes: []Type{}, ReturnValueType: Int64Kind}, stringLength)

	return &repo
}

func initArrayBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register(
		"size",
		FunctionObjectType{ParameterTypes: []Type{}, ReturnValueType: Int64Kind},
		arraySize,
	)
	repo.register(
		"push",
		FunctionObjectType{
			ParameterTypes:  []Type{Int64Kind},
			ReturnValueType: ArrayKind,
		},
		arrayPush,
	)
	repo.register(
		"pushMultiple",
		FunctionObjectType{
			ParameterTypes:  []Type{AnyKind, Int64Kind},
			ReturnValueType: ArrayKind,
		},
		arrayPushMultiple,
	)
	repo.register(
		"delete",
		FunctionObjectType{
			ParameterTypes:  []Type{Int64Kind, Int64Kind},
			ReturnValueType: ArrayKind,
		},
		arrayDelete,
	)
	repo.register(
		"slice",
		FunctionObjectType{
			ParameterTypes:  []Type{Int64Kind, Int64Kind},
			ReturnValueType: ArrayKind,
		},
		arraySlice,
	)
	return &repo
}

func numberToString(params ...Object) Object {
	number, _ := params[0].(*Number)
	return &String{number.Inspect()}
}

func stringLength(params ...Object) Object {
	str, _ := params[0].(*String)
	return &Number{Value: uint64(len(str.Value)), Kind: UInt64Kind}
}

func arraySize(params ...Object) Object {
	arr, _ := params[0].(*Array)
	return &Number{Value: uint64(len(arr.Elements)), Kind: UInt64Kind}
}

func arrayPush(params ...Object) Object {
	arr, _ := params[0].(*Array)
	elem := params[1]

	arr.Elements = append(arr.Elements, elem)
	return arr
}

func arrayDelete(params ...Object) Object {
	arr, _ := params[0].(*Array)
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
	arr, _ := params[0].(*Array)
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
	arr, _ := params[0].(*Array)
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
