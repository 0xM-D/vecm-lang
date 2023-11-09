package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalNewExpression(exp *ast.NewExpression,
	env *object.Environment) object.Object {

	_, isArray := exp.Type.(ast.ArrayType)
	if isArray {
		return evalNewArrayExpression(exp, env)
	}

	_, isHash := exp.Type.(ast.HashType)
	if isHash {
		return evalNewHashExpression(exp, env)
	}

	return newError("New operator not yet supported on type: %s", exp.Type.String())

}

func evalNewArrayExpression(exp *ast.NewExpression, env *object.Environment) object.Object {
	elements := evalExpressionsArray(exp.InitializationList, env)

	arrayType := exp.Type.(ast.ArrayType)

	elementType, _ := evalType(arrayType.ElementType, env)

	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements, ArrayObjectType: object.ArrayObjectType{ElementType: elementType}}
}

func evalNewHashExpression(exp *ast.NewExpression, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for _, expr := range exp.InitializationList {
		pair, ok := expr.(*ast.PairExpression)

		if !ok {
			return newError("Found non pair element in hash initialization list")
		}

		key := object.UnwrapReferenceObject(Eval(pair.Left, env))
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type().Signature())
		}

		value := Eval(pair.Right, env)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs, HashObjectType: object.HashObjectType{KeyType: object.AnyKind, ValueType: object.AnyKind}}
}
