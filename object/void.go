package object

type VoidObjectType struct{}

func (n VoidObjectType) Signature() string { return "void" }

type Void struct{}

func (n *Void) Type() ObjectType { return VOID_OBJ() }
func (n *Void) Inspect() string  { return "void" }
