package object

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ErrorKind }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
