package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) (object.Object, error) {
	if node.ReturnValue == nil {
		return &object.ReturnValue{Value: nil, ReturnValueObjectType: object.ReturnValueObjectType{ReturnType: object.VoidKind}}, nil
	}

	val, err := r.Eval(node.ReturnValue, env)
	if err != nil {
		return nil, err
	}
	return &object.ReturnValue{Value: val, ReturnValueObjectType: object.ReturnValueObjectType{ReturnType: val.Type()}}, nil
}
