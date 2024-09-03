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
		return r.evalArrayIndexExpression(left.(*object.Array), index)
	case object.IsHash(left):
		return r.evalHashIndexExpression(left.(*object.Hash), index)
	default:
		return nil, fmt.Errorf("index operator not supported: %s", left.Type().Signature())
	}
}

func (r *Runtime) evalArrayIndexExpression(array *object.Array, index object.Object) (object.Object, error) {
	idxObj, err := typeCast(index, object.Int64Kind, EXPLICIT_CAST)
	if err != nil {
		return nil, err
	}
	idx := idxObj.(*object.Number).GetInt64()

	maxIndex := int64(len(array.Elements) - 1)

	if idx < 0 || idx > maxIndex {
		return NULL, nil
	}

	return &object.ArrayElementReference{
		Array: array,
		Index: idx,
		ReferenceType: object.ReferenceType{
			IsConstantReference: array.IsConstant(),
			ValueType:           array.ElementType,
		},
	}, nil
}

func (r *Runtime) evalHashIndexExpression(hash *object.Hash, index object.Object) (object.Object, error) {
	key, ok := index.(object.Hashable)
	if !ok {
		return nil, fmt.Errorf("unusable as hash key: %s", index.Type().Signature())
	}

	_, exists := hash.Pairs[key.HashKey()]
	if !exists {
		return NULL, nil
	}

	return &object.HashElementReference{
		Hash: hash,
		Key:  index,
		ReferenceType: object.ReferenceType{
			IsConstantReference: hash.IsConstant(),
			ValueType:           hash.ValueType,
		},
	}, nil
}
