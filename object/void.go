package object

type Void struct{}

func (n *Void) Type() ObjectType { return VoidKind }
func (n *Void) Inspect() string  { return VoidKind.Signature() }
