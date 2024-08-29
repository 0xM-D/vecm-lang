package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalStringLiteral(node *ast.StringLiteral) (object.Object, error) {
	return &object.String{Value: node.Value}, nil
}
