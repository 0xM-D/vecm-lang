package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalDeclarationStatement(declNode *ast.DeclarationStatement, env *object.Environment) object.Object {
	val := object.UnwrapReferenceObject(Eval(declNode.Value, env))
	var expectedType object.ObjectType

	if declNode.Type != nil {
		var err error
		expectedType, err = evalType(declNode.Type, env)

		if err != nil {
			return newError(err.Error())
		}

	}

	if object.IsError(val) {
		return val
	}

	if object.IsNumber(val) && expectedType == nil {
		expectedType = object.Int64Kind
	}

	if expectedType != nil {
		cast := typeCast(val, expectedType, EXPLICIT_CAST)
		if !object.IsError(cast) {
			val = cast
		}
	}

	if object.IsError(val) {
		return val
	}

	if expectedType != nil && !object.TypesMatch(expectedType, val.Type()) {
		return newError("Expression of type %s cannot be assigned to %s", val.Type().Signature(), expectedType.Signature())
	}

	newObject := env.Declare(declNode.Name.Value, declNode.IsConstant, val)

	if newObject == nil {
		return newError("Identifier with name %s already exists.", declNode.Name.Value)
	}

	return newObject
}
