package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalExplicitTypeCast(node *ast.TypeCastExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if object.IsError(left) {
		return left
	}

	castToType, error := evalType(node.Type, env)
	if error != nil {
		return newError(error.Error())
	}

	casted := typeCast(left, castToType, EXPLICIT_CAST)
	if object.IsError(casted) {
		return casted
	}

	return casted
}
