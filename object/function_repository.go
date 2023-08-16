package object

type GenericBuiltinFunction func(...Object) Object

type BuiltinFunction struct {
	BoundParams []Object
	Name        string
	FunctionObjectType
	Function GenericBuiltinFunction
}

func (e BuiltinFunction) Inspect() string {
	return e.FunctionObjectType.Signature() + " " + e.Name
}

func (e BuiltinFunction) Type() ObjectType { return BuiltinFunctionKind }

type FunctionRepository struct {
	Functions map[string]*BuiltinFunction
}

func (fr FunctionRepository) register(name string, functionType FunctionObjectType, function GenericBuiltinFunction) {
	fr.Functions[name] = &BuiltinFunction{BoundParams: []Object{}, Name: name, FunctionObjectType: functionType, Function: function}
}

func (fr FunctionRepository) Get(name string) *BuiltinFunction {
	return fr.Functions[name]
}
