package object

type Environment struct {
	typeStore map[string]ObjectType
	store     map[string]*ObjectReference
	outer     *Environment
}

var GLOBAL_TYPES = map[ObjectKind]bool{
	IntegerKind: true,
	BooleanKind: true,
	NullKind:    true,
	StringKind:  true,
	VoidKind:    true,
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
	_, exists := GLOBAL_TYPES[ObjectKind(name)]
	if exists {
		return ObjectKind(name), true
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
