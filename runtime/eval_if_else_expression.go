package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalIfExpression(ie *ast.IfStatement, env *object.Environment) (object.Object, error) {
	condition, err := r.Eval(ie.Condition, env)
	if err != nil {
		return nil, err
	}
	switch {
	case isTruthy(condition):
		return r.Eval(ie.Consequence, env)
	case ie.Alternative != nil:
		return r.Eval(ie.Alternative, env)
	default:
		return NULL, nil
	}
}
