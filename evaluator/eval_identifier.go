package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	reference := env.GetReference(node.Value)
	if reference == nil {
		return newError("identifier not found: " + node.Value)
	}
	return reference
}
