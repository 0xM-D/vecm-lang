package object

type NullObjectType struct{}

func (n *NullObjectType) Signature() string { return "null" }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ() }
func (n *Null) Inspect() string  { return "null" }
