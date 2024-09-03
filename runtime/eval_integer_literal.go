package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalIntegerLiteral(node *ast.IntegerLiteral, _ *object.Environment) (object.Object, error) {
	kind, err := getMinimumIntegerType(&node.Value)
	if err != nil {
		return nil, err
	}

	if !object.IS_SIGNED[kind] {
		return &object.Number{Value: node.Value.Uint64(), Kind: object.UInt64Kind}, nil
	}

	return &object.Number{Value: object.Int64Bits(node.Value.Int64()), Kind: object.Int64Kind}, nil
}
