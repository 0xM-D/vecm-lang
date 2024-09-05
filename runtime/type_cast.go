package runtime

import (
	"fmt"
	"math"

	"github.com/DustTheory/interpreter/object"
)

type CastType = int

const (
	_ int = iota
	IMPLICIT_CAST
	EXPLICIT_CAST
)

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

func typeCast(obj object.Object, targetType object.ObjectType, castType CastType) (object.Object, error) {
	if obj.Type().Signature() == targetType.Signature() {
		return obj, nil
	}

	if object.IsArray(obj) && targetType.Kind() == object.ArrayKind {
		return arrayCast(obj.(*object.Array), targetType.(*object.ArrayObjectType), castType)
	}

	if object.IsNumber(obj) && object.IsNumberKind(targetType.Kind()) {
		casted, err := numberCast(obj.(*object.Number), targetType.Kind(), castType)
		if err != nil {
			return nil, err
		}
		return casted, nil
	}

	if object.IsNumber(obj) && targetType.Kind() == object.StringKind {
		return &object.String{Value: obj.Inspect()}, nil
	}

	return nil, fmt.Errorf("type cast from %s to %s is not defined", obj.Type().Signature(), targetType.Signature())
}

func numberCast(number *object.Number, target object.ObjectKind, castType CastType) (*object.Number, error) {
	if number.Kind == target {
		return number, nil
	}

	numberWeight := numberCastWeight[number.Type().Kind()]
	targetWeight := numberCastWeight[target]

	if numberWeight > targetWeight && castType == IMPLICIT_CAST {
		return nil, fmt.Errorf("cannot implicitly cast %s into %s", number.Type().Kind(), target.Kind())
	}
	var value uint64
	switch {
	case object.IsInteger(number) && !object.IsIntegerKind(target):
		// Casting from int to float
		switch {
		case number.IsSigned() && target.Kind() == object.Float64Kind:
			value = math.Float64bits(float64(number.GetInt64()))
		case number.IsSigned() && target.Kind() == object.Float32Kind:
			value = uint64(math.Float32bits(float32(number.GetInt64())))
		case number.IsUnsigned() && target.Kind() == object.Float64Kind:
			value = math.Float64bits(float64(number.GetUInt64()))
		case number.IsUnsigned() && target.Kind() == object.Float32Kind:
			value = uint64(math.Float32bits(float32(number.GetUInt64())))
		}
	case object.IsFloat(number) && object.IsIntegerKind(target):
		// casting from float to int
		switch {
		case object.IS_SIGNED[target] && object.IsFloat32(number):
			value = uint64(int64(number.GetFloat32()))
		case object.IS_SIGNED[target] && object.IsFloat64(number):
			value = uint64(int64(number.GetFloat64()))
		case !object.IS_SIGNED[target] && object.IsFloat32(number):
			value = uint64(number.GetFloat32())
		case !object.IS_SIGNED[target] && object.IsFloat64(number):
			value = uint64(number.GetFloat64())
		}
	case object.IsFloat(number) && (target == object.Float32Kind || target == object.Float64Kind):
		// casting from float to float
		if number.Type() == object.Float32Kind && target == object.Float64Kind {
			value = math.Float64bits(float64(number.GetFloat32()))
		} else if number.Type() == object.Float64Kind && target == object.Float32Kind {
			value = uint64(math.Float32bits(float32(number.GetFloat64())))
		}
	default:
		// casting from int to int
		switch target {
		case object.Int8Kind:
			value = object.Int64Bits(int64(int8(number.GetInt64())))
		case object.Int16Kind:
			value = object.Int64Bits(int64(int16(number.GetInt64())))
		case object.Int32Kind:
			value = object.Int64Bits(int64(int32(number.GetInt64())))
		case object.Int64Kind:
			value = object.Int64Bits(number.GetInt64())
		case object.UInt8Kind:
			value = uint64(uint8(number.GetInt64()))
		case object.UInt16Kind:
			value = uint64(uint16(number.GetInt64()))
		case object.UInt32Kind:
			value = uint64(uint32(number.GetInt64()))
		case object.UInt64Kind:
			value = uint64(number.GetInt64())
		}
	}

	return &object.Number{Value: value, Kind: target}, nil
}

func arrayCast(array *object.Array, targetType *object.ArrayObjectType, castType CastType) (object.Object, error) {
	targetElementKind := targetType.ElementType.Kind()

	if object.IsNumberKind(array.ElementType.Kind()) && !object.IsNumberKind(targetElementKind) {
		return nil, fmt.Errorf("type cast from %s to %s is not defined", array.Type().Signature(), targetType.Signature())
	}

	newArray := &object.Array{ArrayObjectType: *targetType, Elements: []object.Object{}}

	for _, number := range array.Elements {
		casted, err := typeCast(number, targetElementKind, castType)
		if err != nil {
			return nil, err
		}
		newArray.Elements = append(newArray.Elements, casted)
	}

	return newArray, nil
}

func arithmeticCast(first, second *object.Number) (*object.Number, *object.Number, error) {
	if first.Type().Kind() == second.Type().Kind() {
		return first, second, nil
	}

	firstWeight := numberCastWeight[first.Type().Kind()]
	secondWeight := numberCastWeight[second.Type().Kind()]

	if firstWeight < secondWeight {
		castedFirst, err := numberCast(first, second.Kind, EXPLICIT_CAST)
		if err != nil {
			return nil, nil, err
		}
		return castedFirst, second, nil
	} else {
		castedSecond, err := numberCast(second, first.Kind, EXPLICIT_CAST)
		if err != nil {
			return nil, nil, err
		}
		return first, castedSecond, nil
	}
}
