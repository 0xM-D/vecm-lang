package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Hashable interface {
	HashKey() HashKey
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
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
func (h *Hash) Type() ObjectType { return HASH_OBJ }
