package object

import (
	"bytes"
	"strings"
)

type ArrayObjectType struct {
	ElementType ObjectType
}

func (a *ArrayObjectType) Signature() string             { return a.ElementType.Signature() + "[]" }
func (a *ArrayObjectType) Kind() ObjectKind              { return ArrayKind }
func (a *ArrayObjectType) Builtins() *FunctionRepository { return ArrayKind.Builtins() }
func (a *ArrayObjectType) IsConstant() bool              { return false }

type Array struct {
	ArrayObjectType
	Elements []Object
}

func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
func (a *Array) Type() ObjectType { return &a.ArrayObjectType }
