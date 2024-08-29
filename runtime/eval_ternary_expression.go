package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalTernaryExpression(node *ast.TernaryExpression, env *object.Environment) (object.Object, error) {
	conditionResult, err := r.Eval(node.Condition, env)
	if err != nil {
		return nil, err
	}

	var expressionToEvaluate ast.Node
	if isTruthy(conditionResult) {
		expressionToEvaluate = node.ValueIfTrue
	} else {
		expressionToEvaluate = node.ValueIfFalse
	}

	result, err := r.Eval(expressionToEvaluate, env)
	if err != nil {
		return nil, err
	}

	return result, nil
}
