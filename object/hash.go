package object

import (
	"bytes"
	"fmt"
	"strings"
)

type HashObjectType struct {
	KeyType   ObjectType
	ValueType ObjectType
}

func (h *HashObjectType) Signature() string {
	var out bytes.Buffer

	out.WriteString("{")
	out.WriteString(h.KeyType.Signature())
	out.WriteString(" -> ")
	out.WriteString(h.ValueType.Signature())
	out.WriteString("}")

	return out.String()
}

func (h *HashObjectType) Kind() ObjectKind              { return HashKind }
func (h *HashObjectType) Builtins() *FunctionRepository { return FunctionKind.Builtins() }
func (h *HashObjectType) IsConstant() bool              { return false }

type Hashable interface {
	HashKey() HashKey
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	HashObjectType
	Pairs map[HashKey]HashPair
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	kvpairs := []string{}
	for _, pair := range h.Pairs {
		kvpairs = append(kvpairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(kvpairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (h *Hash) Type() ObjectType { return &h.HashObjectType }
