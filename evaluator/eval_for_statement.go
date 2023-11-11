package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalForStatement(node *ast.ForStatement, env *object.Environment) (object.Object, error) {
	forEnv := object.NewEnclosedEnvironment(env)

	if node.Initialization != nil {
		_, err := Eval(node.Initialization, forEnv)
		if err != nil {
			return nil, err
		}
	}

	for {
		if node.Condition != nil {
			conditionResult, err := Eval(node.Condition, forEnv)

			if err != nil {
				return nil, err
			}

			if !isTruthy(conditionResult) {
				break
			}
		}

		if node.Body != nil {
			_, err := Eval(node.Body, forEnv)

			if err != nil {
				return nil, err
			}
		}

		if node.AfterThought != nil {
			_, err := Eval(node.AfterThought, forEnv)

			if err != nil {
				return nil, err
			}
		}

	}

	return nil, nil
}
