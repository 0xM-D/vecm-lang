package object

type StringObjectType struct{}

func (s *StringObjectType) Signature() string { return "string" }

type String struct {
	Value string
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ() }
