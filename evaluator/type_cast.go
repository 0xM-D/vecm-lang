package evaluator

import (
	"strconv"

	"github.com/0xM-D/interpreter/object"
)

const (
	IMPLICIT_CAST = true
	EXPLICIT_CAST = false
)

type CastRuleSignature struct {
	from string
	to   string
}

type CastRule struct {
	allowImplicit bool
	cast          func(object.Object) object.Object
}

var castRules = map[CastRuleSignature]CastRule{
	{"int", "string"}:                         {false, intToString},
	{"{string -> string}", "{int -> string}"}: {true, castEmptyMap},
	{"int[]", "string[]"}:                     {true, castEmptyArray},
	{"any[]", "int[]"}:                        {true, castEmptyArray},
}

func typeCast(obj object.Object, targetType object.ObjectType, implicit bool) object.Object {
	fromSignature := obj.Type().Signature()
	toSignature := targetType.Signature()
	castRuleSignature := CastRuleSignature{fromSignature, toSignature}

	castRule, castRuleExists := castRules[castRuleSignature]
	if !castRuleExists {
		return newError("No rule to cast between %s and %s is defined", fromSignature, toSignature)
	}

	if implicit && !castRule.allowImplicit {
		return newError("Implicit cast not allowed between %s and %s", fromSignature, toSignature)
	}

	return castRule.cast(obj)
}

func intToString(obj object.Object) object.Object {
	integer := object.UnwrapReferenceObject(obj).(*object.Integer)
	return &object.String{Value: strconv.FormatInt(integer.Value, 10)}
}

func castEmptyMap(obj object.Object) object.Object {
	hash := object.UnwrapReferenceObject(obj).(*object.Hash)
	if len(hash.Pairs) != 0 {
		return newError("Can't cast non empty hash")
	}
	return &object.Hash{Pairs: hash.Pairs, HashObjectType: object.HashObjectType{KeyType: object.IntegerKind, ValueType: object.IntegerKind}}
}

func castEmptyArray(obj object.Object) object.Object {
	array := object.UnwrapReferenceObject(obj).(*object.Array)
	if len(array.Elements) != 0 {
		return newError("Can't cast non empty array")
	}
	return &object.Array{Elements: array.Elements, ArrayObjectType: object.ArrayObjectType{ElementType: object.StringKind}}
}
