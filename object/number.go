package object

import (
	"fmt"
)

type Number[V int64 | float32 | float64] struct {
	Value V
}

func (i Number[V]) Inspect() string {
	switch val := any(i.Value).(type) {
	case int64:
		return fmt.Sprintf("%d", val)
	case float32:
		return fmt.Sprintf("%gf", val)
	case float64:
		return fmt.Sprintf("%g", val)
	}
	return "This should never ever be reached"
}

func (i Number[V]) Type() ObjectType {
	switch any(i.Value).(type) {
	case int64:
		return IntegerKind
	case float32:
		return Float32Kind
	case float64:
		return Float64Kind
	}
	return ErrorKind
}

func NewNumberOfKind(kind ObjectKind) Object {
	switch kind {
	case IntegerKind:
		return Number[int64]{}
	case Float32Kind:
		return Number[float32]{}
	case Float64Kind:
		return Number[float64]{}
	}
	return &Error{Message: fmt.Sprintf("%s is not a number", kind)}
}
