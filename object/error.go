package object

type ErrorObjectType struct{}

func (e *ErrorObjectType) Signature() string { return "error" }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ() }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
