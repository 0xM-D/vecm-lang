package object

import "fmt"

type IntegerObjectType struct{}

func (i *IntegerObjectType) Signature() string { return "int" }

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ() }
