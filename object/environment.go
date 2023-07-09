package object

type Environment struct {
	typeStore map[string]ObjectType
	store     map[string]Object
	outer     *Environment
}

var GLOBAL_TYPES = map[string]bool{
	INTEGER_OBJ:  true,
	BOOLEAN_OBJ:  true,
	NULL_OBJ:     true,
	FUNCTION_OBJ: true,
	STRING_OBJ:   true,
	ARRAY_OBJ:    true,
	HASH_OBJ:     true,
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) GetObjectType(name string) (ObjectType, bool) {
	_, ok := GLOBAL_TYPES[name]
	if ok {
		return ObjectType(name), ok
	}
	obj, ok := e.typeStore[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.GetObjectType(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
