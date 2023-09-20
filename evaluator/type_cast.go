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
	{object.Int8Kind.Signature(), object.Int16Kind.Signature()}:  {true, numberCast[int8, int16]},
	{object.Int8Kind.Signature(), object.Int32Kind.Signature()}:  {true, numberCast[int8, int32]},
	{object.Int8Kind.Signature(), object.Int64Kind.Signature()}:  {true, numberCast[int8, int64]},
	{object.Int16Kind.Signature(), object.Int32Kind.Signature()}: {true, numberCast[int16, int32]},
	{object.Int16Kind.Signature(), object.Int64Kind.Signature()}: {true, numberCast[int16, int64]},
	{object.Int32Kind.Signature(), object.Int64Kind.Signature()}: {true, numberCast[int32, int64]},

	{object.Int8Kind.Signature(), object.Float32Kind.Signature()}:  {true, numberCast[int8, float32]},
	{object.Int16Kind.Signature(), object.Float32Kind.Signature()}: {true, numberCast[int16, float32]},
	{object.Int32Kind.Signature(), object.Float32Kind.Signature()}: {true, numberCast[int32, float32]},
	{object.Int64Kind.Signature(), object.Float32Kind.Signature()}: {true, numberCast[int64, float32]},

	{object.Int8Kind.Signature(), object.Float64Kind.Signature()}:  {true, numberCast[int8, float64]},
	{object.Int16Kind.Signature(), object.Float64Kind.Signature()}: {true, numberCast[int16, float64]},
	{object.Int32Kind.Signature(), object.Float64Kind.Signature()}: {true, numberCast[int32, float64]},
	{object.Int64Kind.Signature(), object.Float64Kind.Signature()}: {true, numberCast[int64, float64]},

	{object.Float32Kind.Signature(), object.Float64Kind.Signature()}: {true, numberCast[float32, float64]},
	{"{string -> string}", "{int -> string}"}:                        {true, castEmptyMap},
	{"int[]", "string[]"}:                                            {true, castEmptyArray},
	{"any[]", "int[]"}:                                               {true, castEmptyArray},
}

func typeCast(obj object.Object, targetType object.ObjectType, implicit bool) object.Object {

	fromSignature := obj.Type().Signature()
	toSignature := targetType.Signature()
	castRuleSignature := CastRuleSignature{fromSignature, toSignature}

	if fromSignature == toSignature {
		return obj
	}

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
	integer := object.UnwrapReferenceObject(obj).(*object.Number[int64])
	return &object.String{Value: strconv.FormatInt(integer.Value, 10)}
}

func numberCast[
	F int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64,
	T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](obj object.Object) object.Object {
	val := object.UnwrapReferenceObject(obj).(*object.Number[F])
	return &object.Number[T]{Value: T(val.Value)}
}

func castEmptyMap(obj object.Object) object.Object {
	hash := object.UnwrapReferenceObject(obj).(*object.Hash)
	if len(hash.Pairs) != 0 {
		return newError("Can't cast non empty hash")
	}
	return &object.Hash{Pairs: hash.Pairs, HashObjectType: object.HashObjectType{KeyType: object.Int64Kind, ValueType: object.Int64Kind}}
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
	INT8_WEIGHT
	UINT8_WEIGHT
	INT16_WEIGHT
	UINT16_WEIGHT
	INT32_WEIGHT
	UINT32_WEIGHT
	INT64_WEIGHT
	UINT64_WEIGHT
	FLOAT32_WEIGHT
	FLOAT64_WEIGHT
)

var numberCastWeight = map[object.ObjectKind]uint8{
	object.Int8Kind:    INT8_WEIGHT,
	object.Int16Kind:   INT16_WEIGHT,
	object.Int32Kind:   INT32_WEIGHT,
	object.Int64Kind:   INT64_WEIGHT,
	object.UInt8Kind:   UINT8_WEIGHT,
	object.UInt16Kind:  UINT16_WEIGHT,
	object.UInt32Kind:  UINT32_WEIGHT,
	object.UInt64Kind:  UINT64_WEIGHT,
	object.Float32Kind: FLOAT32_WEIGHT,
	object.Float64Kind: FLOAT64_WEIGHT,
}

func castToLargerNumberType(nums ...object.Object) []object.Object {
	if len(nums) == 0 {
		return []object.Object{}
	}

	kindToCastTo := object.UnwrapReferenceObject(nums[0]).Type().Kind()

	for _, num := range nums {
		currKind := object.UnwrapReferenceObject(num).Type().Kind()
		if numberCastWeight[currKind] > numberCastWeight[kindToCastTo] {
			kindToCastTo = currKind
		}
	}

	resultNums := []object.Object{}

	for _, num := range nums {
		result := typeCast(object.UnwrapReferenceObject(num), kindToCastTo, true)
		if object.IsError(result) {
			println(result.Inspect())
		}
		resultNums = append(resultNums, result)
	}

	return resultNums
}
