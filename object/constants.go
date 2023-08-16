package object

import "strconv"

type ObjectKind string

var intrinsicTypeFunctionRepositories = initIntrinsicTypeBuiltins()

func (o ObjectKind) Signature() string             { return string(o) }
func (o ObjectKind) Kind() ObjectKind              { return o }
func (o ObjectKind) Builtins() *FunctionRepository { return intrinsicTypeFunctionRepositories[o] }

const (
	Invalid             ObjectKind = "invalid"
	IntegerKind         ObjectKind = "int"
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

	repos[IntegerKind] = initIntegerBuiltins()
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

	repo.register("length", FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: IntegerKind}, stringLength)

	return &repo
}

func initArrayBuiltins() *FunctionRepository {
	repo := FunctionRepository{Functions: map[string]*BuiltinFunction{}}

	repo.register("size", FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: IntegerKind}, arraySize)
	repo.register("push", FunctionObjectType{ParameterTypes: []ObjectType{IntegerKind}, ReturnValueType: IntegerKind}, arrayPush)

	return &repo
}

func intToString(params ...Object) Object {
	integer := params[0].(*Integer)
	return &String{strconv.FormatInt(integer.Value, 10)}
}

func stringLength(params ...Object) Object {
	str := params[0].(*String)
	return &Integer{int64(len(str.Value))}
}

func arraySize(params ...Object) Object {
	arr := params[0].(*Array)
	return &Integer{int64(len(arr.Elements))}
}

func arrayPush(params ...Object) Object {
	arr := params[0].(*Array)
	elem := params[1]

	arr.Elements = append(arr.Elements, elem)
	println(len(arr.Elements))
	return arr
}
