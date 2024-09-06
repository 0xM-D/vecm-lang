package runtime

import (
	"fmt"
	"math"

	"github.com/DustTheory/interpreter/object"
)

type CastType = int

const (
	_ int = iota
	ImplicitCast
	ExplicitCast
)

const (
	_ uint8 = iota
	Int8Weight
	UInt8Weight
	Int16Weight
	UInt16Weight
	Int32Weight
	UInt32Weight
	Int64Weight
	UInt64Weight
	Float32Weight
	Float64Weight
)

var numberCastWeight = map[object.Kind]uint8{
	object.Int8Kind:    Int8Weight,
	object.Int16Kind:   Int16Weight,
	object.Int32Kind:   Int32Weight,
	object.Int64Kind:   Int64Weight,
	object.UInt8Kind:   UInt8Weight,
	object.UInt16Kind:  UInt16Weight,
	object.UInt32Kind:  UInt32Weight,
	object.UInt64Kind:  UInt64Weight,
	object.Float32Kind: Float32Weight,
	object.Float64Kind: Float64Weight,
}

func typeCast(obj object.Object, targetType object.Type, castType CastType) (object.Object, error) {
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

func numberCast(number *object.Number, target object.Kind, castType CastType) (*object.Number, error) {
	if number.Kind == target {
		return number, nil
	}

	numberWeight := numberCastWeight[number.Type().Kind()]
	targetWeight := numberCastWeight[target]

	if numberWeight > targetWeight && castType == ImplicitCast {
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
		case object.IsSigned[target] && object.IsFloat32(number):
			value = uint64(int64(number.GetFloat32()))
		case object.IsSigned[target] && object.IsFloat64(number):
			value = uint64(int64(number.GetFloat64()))
		case !object.IsSigned[target] && object.IsFloat32(number):
			value = uint64(number.GetFloat32())
		case !object.IsSigned[target] && object.IsFloat64(number):
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
		//nolint:exhaustive // we are handling all cases
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
		castedFirst, err := numberCast(first, second.Kind, ExplicitCast)
		if err != nil {
			return nil, nil, err
		}
		return castedFirst, second, nil
	}

	castedSecond, err := numberCast(second, first.Kind, ExplicitCast)
	if err != nil {
		return nil, nil, err
	}
	return first, castedSecond, nil
}
