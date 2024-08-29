package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalForStatement(node *ast.ForStatement, env *object.Environment) (object.Object, error) {
	forEnv := object.NewEnclosedEnvironment(env)

	if node.Initialization != nil {
		_, err := r.Eval(node.Initialization, forEnv)
		if err != nil {
			return nil, err
		}
	}

	for {
		if node.Condition != nil {
			conditionResult, err := r.Eval(node.Condition, forEnv)

			if err != nil {
				return nil, err
			}

			if !isTruthy(conditionResult) {
				break
			}
		}

		if node.Body != nil {
			_, err := r.Eval(node.Body, forEnv)

			if err != nil {
				return nil, err
			}
		}

		if node.AfterThought != nil {
			_, err := r.Eval(node.AfterThought, forEnv)

			if err != nil {
				return nil, err
			}
		}

	}

	return nil, nil
}
