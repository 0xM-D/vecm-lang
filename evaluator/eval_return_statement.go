package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) (object.Object, error) {
	val, err := Eval(node.ReturnValue, env)
	if err != nil {
		return nil, err
	}
	return &object.ReturnValue{Value: val, ReturnValueObjectType: object.ReturnValueObjectType{ReturnType: val.Type()}}, nil
}
