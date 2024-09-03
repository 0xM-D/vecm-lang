package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalFunctionDeclarationStatement(
	node *ast.FunctionDeclarationStatement,
	env *object.Environment,
) (object.Object, error) {
	functionType, err := r.evalType(node.Type, env)
	if err != nil {
		return nil, err
	}

	function := &object.Function{
		Parameters:         node.Parameters,
		Env:                env,
		Body:               node.Body,
		FunctionObjectType: *functionType.(*object.FunctionObjectType),
	}

	ref, err := env.Declare(node.Name.Value, true, function)
	if err != nil {
		return nil, fmt.Errorf("error in function declaration: %w", err)
	}
	return ref, nil
}
