package runtime

import "github.com/DustTheory/interpreter/object"

var (
	Null  = &object.Null{}
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
)
