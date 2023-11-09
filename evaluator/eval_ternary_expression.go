package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalTernaryExpression(node *ast.TernaryExpression, env *object.Environment) object.Object {
	conditionResult := Eval(node.Condition, env)
	if object.IsError(conditionResult) {
		return conditionResult
	}

	var expressionToEvaluate ast.Node
	if isTruthy(conditionResult) {
		expressionToEvaluate = node.ValueIfTrue
	} else {
		expressionToEvaluate = node.ValueIfFalse
	}

	result := Eval(expressionToEvaluate, env)
	if object.IsError(result) {
		return result
	}

	return result
}
