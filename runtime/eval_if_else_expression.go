package runtime

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func (r *Runtime) evalIfExpression(ie *ast.IfStatement, env *object.Environment) (object.Object, error) {
	condition, err := r.Eval(ie.Condition, env)
	if err != nil {
		return nil, err
	}
	if isTruthy(condition) {
		return r.Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return r.Eval(ie.Alternative, env)
	} else {
		return NULL, nil
	}
}
