package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	functionType, err := evalType(node.Type, env)
	if err != nil {
		return newError(err.Error())
	}

	function := &object.Function{
		Parameters:         node.Parameters,
		Env:                env,
		Body:               node.Body,
		FunctionObjectType: *functionType.(*object.FunctionObjectType),
	}
	return function
}
