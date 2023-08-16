package object

type Null struct{}

func (n *Null) Type() ObjectType { return NullKind }
func (n *Null) Inspect() string  { return NullKind.Signature() }
