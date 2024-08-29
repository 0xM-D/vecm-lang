package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) (object.Object, error) {
	reference := env.GetReference(node.Value)
	if reference == nil {
		return nil, fmt.Errorf("identifier not found: " + node.Value)
	}
	return reference, nil
}
