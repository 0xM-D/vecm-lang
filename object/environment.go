package object

import "fmt"

type EnvStoreEntry struct {
	Object
	IsConstant bool
}

type Environment struct {
	typeStore map[string]ObjectType
	store     map[string]*EnvStoreEntry
	outer     *Environment
}

var GLOBAL_TYPES = map[ObjectKind]ObjectKind{
	"char":      Int8Kind,
	"int":       Int64Kind,
	Int8Kind:    Int8Kind,
	Int16Kind:   Int16Kind,
	Int32Kind:   Int32Kind,
	Int64Kind:   Int64Kind,
	UInt8Kind:   UInt8Kind,
	UInt16Kind:  UInt16Kind,
	UInt32Kind:  UInt32Kind,
	UInt64Kind:  Int64Kind,
	Float32Kind: Float32Kind,
	Float64Kind: Float64Kind,
	BooleanKind: BooleanKind,
	NullKind:    NullKind,
	StringKind:  StringKind,
	VoidKind:    VoidKind,
}

func NewEnvironment() *Environment {
	s := make(map[string]*EnvStoreEntry)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) GetReference(name string) Object {
	entry, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.GetReference(name)
	}
	if !ok {
		return nil
	}
	return &VariableReference{Env: e, Name: name, ReferenceType: ReferenceType{IsConstantReference: entry.IsConstant, ValueType: entry.Object.Type()}}
}

func (e *Environment) Get(name string) Object {
	entry, ok := e.store[name]

	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	if !ok {
		return nil
	}
	return entry.Object
}

func (e *Environment) GetObjectType(name string) (ObjectType, bool) {
	globalObjectType, globalObjectTypeExists := GLOBAL_TYPES[ObjectKind(name)]
	if globalObjectTypeExists {
		return globalObjectType, true
	}
	objectType, objectTypeExists := e.typeStore[name]
	if !objectTypeExists && e.outer != nil {
		objectType, objectTypeExists = e.outer.GetObjectType(name)
	}
	return objectType, objectTypeExists
}

func (e *Environment) Declare(name string, isConstant bool, val Object) ObjectReference {
	_, exists := e.store[name]
	if exists {
		return nil
	}
	newReference := &VariableReference{e, name, ReferenceType{isConstant, val.Type()}}
	e.store[name] = &EnvStoreEntry{val, isConstant}
	return newReference
}

func (e *Environment) Set(name string, val Object) (Object, error) {
	entry, exists := e.store[name]
	if exists && entry.IsConstant {
		return nil, fmt.Errorf("Cannot assign to const variable")
	}
	e.store[name] = &EnvStoreEntry{val, entry.IsConstant}
	return e.store[name].Object, nil
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
