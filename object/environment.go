package object

type Environment struct {
	typeStore map[string]ObjectType
	store     map[string]*Object
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

func (e *Environment) Declare(name string, isConstant bool, val ObjectValue) *Object {
	obj := e.store[name]
	if obj != nil {
		return nil
	}
	e.store[name] = &Object{IsConstant: isConstant, Value: val}
	return e.store[name]
}

func (e *Environment) Set(name string, val ObjectValue) *Object {
	entry := e.store[name]
	if entry != nil {
		if entry.IsConstant {
			return nil
		}
		entry.Value = val
	}
	return entry
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
