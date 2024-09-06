package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalExplicitTypeCast(node *ast.TypeCastExpression, env *object.Environment) (object.Object, error) {
	left, err := r.Eval(node.Left, env)
	if err != nil {
		return nil, err
	}

	castToType, err := r.evalType(node.Type, env)
	if err != nil {
		return nil, err
	}

	casted, err := typeCast(left, castToType, ExplicitCast)
	if err != nil {
		return nil, err
	}

	return casted, nil
}
