package object

import (
	"fmt"
)

type Number[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64] struct {
	Value V
}

func (i *Number[V]) Inspect() string {
	switch val := any(i.Value).(type) {
	case float32:
		return fmt.Sprintf("%gf", val)
	case float64:
		return fmt.Sprintf("%g", val)
	default:
		return fmt.Sprintf("%d", val)
	}
}

func (i *Number[V]) Type() ObjectType {
	switch any(i.Value).(type) {
	case int8:
		return Int8Kind
	case int16:
		return Int16Kind
	case int32:
		return Int32Kind
	case int64:
		return Int64Kind
	case uint8:
		return UInt8Kind
	case uint16:
		return UInt16Kind
	case uint32:
		return UInt32Kind
	case uint64:
		return UInt64Kind
	case float32:
		return Float32Kind
	case float64:
		return Float64Kind
	}
	return ErrorKind
}

func NewNumberOfKind(kind ObjectKind) Object {
	switch kind {
	case Int64Kind:
		return &Number[int64]{}
	case Float32Kind:
		return &Number[float32]{}
	case Float64Kind:
		return &Number[float64]{}
	}
	return &Error{Message: fmt.Sprintf("%s is not a number", kind)}
}
