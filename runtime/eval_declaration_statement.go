package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalDeclarationStatement(
	declNode *ast.DeclarationStatement,
	env *object.Environment,
) (object.Object, error) {
	ref, err := r.Eval(declNode.Value, env)
	var expectedType object.ObjectType

	if err != nil {
		return nil, fmt.Errorf("error returned from external package: %w", err)
	}

	if declNode.Type != nil {
		var typeErr error
		expectedType, typeErr = r.evalType(declNode.Type, env)
		if typeErr != nil {
			return nil, typeErr
		}
	}

	val := object.UnwrapReferenceObject(ref)

	if object.IsNumber(val) && expectedType == nil {
		expectedType = object.Int64Kind
	}

	if expectedType != nil {
		cast, castErr := typeCast(val, expectedType, EXPLICIT_CAST)
		if castErr == nil {
			val = cast
		}
	}

	if expectedType != nil && !object.TypesMatch(expectedType, val.Type()) {
		return nil, fmt.Errorf("expression of type %s cannot be assigned to %s",
			val.Type().Signature(), expectedType.Signature())
	}

	newObject, err := env.Declare(declNode.Name.Value, declNode.IsConstant, val)
	if err != nil {
		return nil, fmt.Errorf("failed to declare variable %s: %w", declNode.Name.Value, err)
	}

	return newObject, nil
}
