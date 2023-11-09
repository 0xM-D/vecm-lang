package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalIndexExpression(node *ast.IndexExpression, env *object.Environment) object.Object {
	left := object.UnwrapReferenceObject(Eval(node.Left, env))
	if object.IsError(left) {
		return left
	}
	index := object.UnwrapReferenceObject(Eval(node.Index, env))
	if object.IsError(index) {
		return index
	}
	switch {
	case object.IsArray(left) && object.IsInteger(index):
		return evalArrayIndexExpression(left, index)
	case object.IsHash(left):
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type().Signature())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)

	idx := typeCast(index, object.Int64Kind, true).(*object.Number).GetInt64()
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return &object.ArrayElementReference{Array: arrayObject, Index: idx}
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type().Signature())
	}

	_, exists := hashObject.Pairs[key.HashKey()]
	if !exists {
		return NULL
	}

	return &object.HashElementReference{Hash: hashObject, Key: index}
}
