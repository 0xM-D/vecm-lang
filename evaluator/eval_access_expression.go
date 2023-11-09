package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalAccessExpression(node *ast.AccessExpression, env *object.Environment) object.Object {
	left := object.UnwrapReferenceObject(Eval(node.Left, env))
	if object.IsError(left) {
		return left
	}
	right, ok := node.Right.(*ast.Identifier)
	if !ok {
		return newError("Right side of access expression is not an identifier")
	}

	repo := left.Type().Builtins()
	var member *object.BuiltinFunction
	if repo != nil {
		member = left.Type().Builtins().Get(right.Value)
	}
	if member == nil {
		return newError("Member %s does not exist on %s", right, left.Type().Signature())
	}

	if object.IsBuiltinFunction(member) {
		return object.BuiltinFunction{BoundParams: []object.Object{left}, Function: member.Function, FunctionObjectType: member.FunctionObjectType, Name: member.Name}
	}
	return member
}
