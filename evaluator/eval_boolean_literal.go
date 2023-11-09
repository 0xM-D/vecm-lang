package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalBooleanLiteral(node *ast.BooleanLiteral) object.Object {
	return nativeBoolToBooleanObject(node.Value)
}
