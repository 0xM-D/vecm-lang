package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalIndexExpression(node *ast.IndexExpression, env *object.Environment) (object.Object, error) {
	left, err := r.Eval(node.Left, env)
	if err != nil {
		return nil, err
	}
	index, err := r.Eval(node.Index, env)
	if err != nil {
		return nil, err
	}

	left, index = object.UnwrapReferenceObject(left), object.UnwrapReferenceObject(index)

	switch {
	case object.IsArray(left) && object.IsInteger(index):
		return r.evalArrayIndexExpression(left, index)
	case object.IsHash(left):
		return r.evalHashIndexExpression(left, index)
	default:
		return nil, fmt.Errorf("index operator not supported: %s", left.Type().Signature())
	}
}

func (r *Runtime) evalArrayIndexExpression(array, index object.Object) (object.Object, error) {
	arrayObject := array.(*object.Array)

	idxObj, err := typeCast(index, object.Int64Kind, EXPLICIT_CAST)
	if err != nil {
		return nil, err
	}
	idx := idxObj.(*object.Number).GetInt64()

	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL, nil
	}

	return &object.ArrayElementReference{Array: arrayObject, Index: idx}, nil
}

func (r *Runtime) evalHashIndexExpression(hash, index object.Object) (object.Object, error) {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return nil, fmt.Errorf("unusable as hash key: %s", index.Type().Signature())
	}

	_, exists := hashObject.Pairs[key.HashKey()]
	if !exists {
		return NULL, nil
	}

	return &object.HashElementReference{Hash: hashObject, Key: index}, nil
}
