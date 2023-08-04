package object

type Environment struct {
	typeStore map[string]ObjectType
	store     map[string]*Object
	outer     *Environment
}

var GLOBAL_TYPES = map[string]ObjectType{
	"int":    INTEGER_OBJ(),
	"bool":   BOOLEAN_OBJ(),
	"null":   NULL_OBJ(),
	"fn":     &FunctionObjectType{ParameterTypes: []ObjectType{}, ReturnValueType: INTEGER_OBJ()},
	"string": STRING_OBJ(),
	"array":  &ArrayObjectType{ElementType: INTEGER_OBJ()},
	"hash":   &HashObjectType{KeyType: INTEGER_OBJ(), ValueType: INTEGER_OBJ()},
}

func NewEnvironment() *Environment {
	s := make(map[string]*Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) *Object {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj = e.outer.Get(name)
	}
	return obj
}

func (e *Environment) GetObjectType(name string) (ObjectType, bool) {
	objectType, ok := GLOBAL_TYPES[name]
	if ok {
		return objectType, true
	}
	obj, ok := e.typeStore[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.GetObjectType(name)
	}
	return obj, ok
}

func (e *Environment) Declare(name string, isConstant bool, val Object) *Object {
	obj := e.store[name]
	if obj != nil {
		return nil
	}
	e.store[name] = &val
	return e.store[name]
}

func (e *Environment) Set(name string, val Object) Object {
	entry := e.store[name]
	if entry != nil {
		// if entry.IsConstant() {
		// 	return nil
		// }
		entry = &val
	}
	return *entry
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
