package object

import "fmt"

type BooleanObjectType struct{}

func (b *BooleanObjectType) Signature() string { return "bool" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ() }
