package object

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

type Number struct {
	Value uint64
	Kind  Kind
}

func (n *Number) Inspect() string {
	if IsFloat32(n) {
		return fmt.Sprintf("%gf", n.GetFloat32())
	}
	if IsFloat64(n) {
		return fmt.Sprintf("%g", n.GetFloat64())
	}
	if n.IsSigned() {
		return strconv.FormatInt(n.GetInt64(), 10)
	}
	return strconv.FormatUint(n.GetUInt64(), 10)
}

func (n *Number) Type() Type {
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

var IsSigned = map[Kind]bool{
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
	return IsSigned[n.Kind]
}

func (n *Number) IsUnsigned() bool {
	return !IsSigned[n.Kind]
}

func Int64Bits(x int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&x))
}

func Int64FromBits(x uint64) int64 {
	return *(*int64)(unsafe.Pointer(&x))
}

var NumberTypes = []Kind{
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
