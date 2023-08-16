package object

type AnyObjectType struct{}

func (*AnyObjectType) Signature() string { return "any" }
