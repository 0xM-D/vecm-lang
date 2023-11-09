package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) object.Object {
	val := Eval(node.ReturnValue, env)
	if object.IsError(val) {
		return val
	}
	return &object.ReturnValue{Value: val, ReturnValueObjectType: object.ReturnValueObjectType{ReturnType: val.Type()}}
}
