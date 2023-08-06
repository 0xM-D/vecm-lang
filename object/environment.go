package object

type Environment struct {
	typeStore map[string]ObjectType
	store     map[string]*ObjectReference
	outer     *Environment
}

var GLOBAL_TYPES = map[string]ObjectType{
	"int":    INTEGER_OBJ(),
	"bool":   BOOLEAN_OBJ(),
	"null":   NULL_OBJ(),
	"string": STRING_OBJ(),
}

func NewEnvironment() *Environment {
	s := make(map[string]*ObjectReference)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) *ObjectReference {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
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

func (e *Environment) Declare(name string, isConstant bool, val Object) *ObjectReference {
	_, exists := e.store[name]
	if exists {
		return nil
	}
	newReference := &ObjectReference{val, isConstant, name}
	e.store[name] = newReference
	return newReference
}

func (e *Environment) Set(name string, val Object) *ObjectReference {
	entry, exists := e.store[name]
	if exists {
		if entry.IsConstant {
			return nil
		}
		entry.Object = val
	}
	return entry
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
