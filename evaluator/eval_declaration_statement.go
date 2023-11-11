package evaluator

import (
	"fmt"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalDeclarationStatement(declNode *ast.DeclarationStatement, env *object.Environment) (object.Object, error) {
	ref, err := Eval(declNode.Value, env)
	var expectedType object.ObjectType

	if err != nil {
		return nil, err
	}

	if declNode.Type != nil {
		var err error
		expectedType, err = evalType(declNode.Type, env)

		if err != nil {
			return nil, err
		}

	}

	val := object.UnwrapReferenceObject(ref)

	if object.IsNumber(val) && expectedType == nil {
		expectedType = object.Int64Kind
	}

	if expectedType != nil {
		cast, err := typeCast(val, expectedType, EXPLICIT_CAST)
		if err == nil {
			val = cast
		}
	}

	if expectedType != nil && !object.TypesMatch(expectedType, val.Type()) {
		return nil, fmt.Errorf("expression of type %s cannot be assigned to %s", val.Type().Signature(), expectedType.Signature())
	}

	newObject := env.Declare(declNode.Name.Value, declNode.IsConstant, val)

	if newObject == nil {
		return nil, fmt.Errorf("identifier with name %s already exists", declNode.Name.Value)
	}

	return newObject, nil
}
