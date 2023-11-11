package runtime

import (
	"fmt"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func (r *Runtime) evalAccessExpression(node *ast.AccessExpression, env *object.Environment) (object.Object, error) {
	leftRef, err := r.Eval(node.Left, env)
	if err != nil {
		return nil, err
	}
	right, ok := node.Right.(*ast.Identifier)
	if !ok {
		return nil, fmt.Errorf("right side of access expression is not an identifier")
	}

	left := object.UnwrapReferenceObject(leftRef)
	repo := left.Type().Builtins()

	var member *object.BuiltinFunction
	if repo != nil {
		member = left.Type().Builtins().Get(right.Value)
	}
	if member == nil {
		return nil, fmt.Errorf("member %s does not exist on %s", right, left.Type().Signature())
	}

	if object.IsBuiltinFunction(member) {
		return object.BuiltinFunction{BoundParams: []object.Object{left}, Function: member.Function, FunctionObjectType: member.FunctionObjectType, Name: member.Name}, nil
	}
	return member, nil
}
