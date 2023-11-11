package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalTernaryExpression(node *ast.TernaryExpression, env *object.Environment) (object.Object, error) {
	conditionResult, err := Eval(node.Condition, env)
	if err != nil {
		return nil, err
	}

	var expressionToEvaluate ast.Node
	if isTruthy(conditionResult) {
		expressionToEvaluate = node.ValueIfTrue
	} else {
		expressionToEvaluate = node.ValueIfFalse
	}

	result, err := Eval(expressionToEvaluate, env)
	if err != nil {
		return nil, err
	}

	return result, nil
}
