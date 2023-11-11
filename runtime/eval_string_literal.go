package runtime

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func (r *Runtime) evalStringLiteral(node *ast.StringLiteral) (object.Object, error) {
	return &object.String{Value: node.Value}, nil
}
