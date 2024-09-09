package runtime

import (
	"errors"
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalNewExpression(exp *ast.NewExpression,
	env *object.Environment) (object.Object, error) {
	arrayType, isArray := exp.Type.(ast.ArrayType)
	if isArray {
		return r.evalNewArrayExpression(arrayType, exp.InitializationList, env)
	}

	_, isHash := exp.Type.(ast.HashType)
	if isHash {
		return r.evalNewHashExpression(exp, env)
	}

	return nil, fmt.Errorf("new operator not yet supported on type: %s", exp.Type.String())
}

func (r *Runtime) evalNewArrayExpression(
	arrayType ast.ArrayType,
	initializationList []ast.Expression,
	env *object.Environment,
) (object.Object, error) {
	elements, err := r.evalExpressionsArray(initializationList, env)
	if err != nil {
		return nil, err
	}

	elementType, err := r.evalType(arrayType.ElementType, env)
	if err != nil {
		return nil, err
	}

	return &object.Array{Elements: elements, ArrayObjectType: object.ArrayObjectType{ElementType: elementType}}, nil
}

func (r *Runtime) evalNewHashExpression(exp *ast.NewExpression, env *object.Environment) (object.Object, error) {
	pairs := make(map[object.HashKey]object.HashPair)

	for _, expr := range exp.InitializationList {
		pair, ok := expr.(*ast.PairExpression)

		if !ok {
			return nil, errors.New("found non pair element in hash initialization list")
		}

		key, err := r.Eval(pair.Left, env)
		if err != nil {
			return nil, err
		}

		hashKey, ok := object.UnwrapReferenceObject(key).(object.Hashable)
		if !ok {
			return nil, fmt.Errorf("unusable as hash key: %s", key.Type().Signature())
		}

		value, err := r.Eval(pair.Right, env)
		if err != nil {
			return nil, err
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{
		Pairs: pairs,
		HashObjectType: object.HashObjectType{
			KeyType:   object.AnyKind,
			ValueType: object.AnyKind,
		},
	}, nil
}
