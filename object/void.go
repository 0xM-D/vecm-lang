package object

type Void struct{}

func (n *Void) Type() Type      { return VoidKind }
func (n *Void) Inspect() string { return VoidKind.Signature() }
