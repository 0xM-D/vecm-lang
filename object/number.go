package object

import (
	"fmt"
	"math"
	"unsafe"
)

type Number struct {
	Value uint64
	Kind  ObjectKind
}

func (n *Number) Inspect() string {
	if IsFloat32(n) {
		return fmt.Sprintf("%gf", n.GetFloat32())
	}
	if IsFloat64(n) {
		return fmt.Sprintf("%g", n.GetFloat64())
	}
	if n.IsSigned() {
		return fmt.Sprintf("%d", n.GetInt64())
	}
	return fmt.Sprintf("%d", n.GetUInt64())
}

func (n *Number) Type() ObjectType {
	return n.Kind
}

func (n *Number) GetUInt64() uint64 {
	return n.Value
}

func (n *Number) GetInt64() int64 {
	return Int64FromBits(n.Value)
}

func (n *Number) GetFloat32() float32 {
	return math.Float32frombits(uint32(n.Value))
}

func (n *Number) GetFloat64() float64 {
	return math.Float64frombits(n.Value)
}

var IS_SIGNED = map[ObjectKind]bool{
	Int8Kind:    true,
	Int16Kind:   true,
	Int32Kind:   true,
	Int64Kind:   true,
	Float32Kind: true,
	Float64Kind: true,
	UInt8Kind:   false,
	UInt16Kind:  false,
	UInt32Kind:  false,
	UInt64Kind:  false,
}

func (n *Number) IsSigned() bool {
	return IS_SIGNED[n.Kind]
}

func (n *Number) IsUnsigned() bool {
	return !IS_SIGNED[n.Kind]
}

func Int64Bits(x int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&x))
}

func Int64FromBits(x uint64) int64 {
	return *(*int64)(unsafe.Pointer(&x))
}

var NumberTypes = []ObjectKind{
	Int8Kind,
	Int16Kind,
	Int32Kind,
	Int64Kind,
	UInt8Kind,
	UInt16Kind,
	UInt32Kind,
	UInt64Kind,
	Float32Kind,
	Float64Kind,
}
