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
	{"int", "string"}:                         {true, intToString},
	{"int", "float32"}:                        {true, numberCast[int64, float32]},
	{"int", "float64"}:                        {true, numberCast[int64, float64]},
	{"float32", "float64"}:                    {true, numberCast[float32, float64]},
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
	integer := object.UnwrapReferenceObject(obj).(object.Number[int64])
	return &object.String{Value: strconv.FormatInt(integer.Value, 10)}
}

func numberCast[F int64 | float32 | float64, T int64 | float32 | float64](obj object.Object) object.Object {
	val := object.UnwrapReferenceObject(obj).(object.Number[F])
	return object.Number[T]{Value: T(val.Value)}
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

const (
	_ uint8 = iota
	INT_WEIGHT
	FLOAT32_WEIGHT
	FLOAT64_WEIGHT
)

var numberCastWeight = map[object.ObjectKind]uint8{
	object.IntegerKind: INT_WEIGHT,
	object.Float32Kind: FLOAT32_WEIGHT,
	object.Float64Kind: FLOAT64_WEIGHT,
}

func castToLargerNumberType(num1 object.Object, num2 object.Object) (object.Object, object.Object, object.ObjectKind) {
	num1 = object.UnwrapReferenceObject(num1)
	num2 = object.UnwrapReferenceObject(num2)

	k1 := num1.Type().Kind()
	k2 := num2.Type().Kind()
	w1 := numberCastWeight[k1]
	w2 := numberCastWeight[k2]

	if w1 == w2 {
		return num1, num2, k1
	}
	if w1 > w2 {
		return num1, typeCast(num2, k1, IMPLICIT_CAST), k1
	}
	return typeCast(num1, k2, IMPLICIT_CAST), num2, k2
}
