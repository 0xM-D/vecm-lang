package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
	forEnv := object.NewEnclosedEnvironment(env)

	if node.Initialization != nil {
		initResult := Eval(node.Initialization, forEnv)
		if object.IsError(initResult) {
			return initResult
		}
	}

	for {
		if node.Condition != nil {
			conditionResult := Eval(node.Condition, forEnv)

			if object.IsError(conditionResult) {
				return conditionResult
			}

			if !isTruthy(conditionResult) {
				break
			}
		}

		if node.Body != nil {
			bodyResult := Eval(node.Body, forEnv)

			if bodyResult != nil && object.IsError(bodyResult) {
				return bodyResult
			}
		}

		if node.AfterThought != nil {
			afterThoughtResult := Eval(node.AfterThought, forEnv)

			if object.IsError(afterThoughtResult) {
				return afterThoughtResult
			}
		}

	}

	return nil
}
