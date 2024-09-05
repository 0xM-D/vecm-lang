package object

type Null struct{}

func (n *Null) Type() Type      { return NullKind }
func (n *Null) Inspect() string { return NullKind.Signature() }
