package evaluator

import (
	"math"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
	"github.com/0xM-D/interpreter/token"
)

type OperatorFnSignature struct {
	Operator string
	LType    string
	RType    string
}

type InfixEvalFn func(object.Object, object.Object, *object.Environment) object.Object

var infixEvalFns = map[OperatorFnSignature]InfixEvalFn{
	{"==", string(object.BooleanKind), string(object.BooleanKind)}: booleanEquals,
	{"!=", string(object.BooleanKind), string(object.BooleanKind)}: booleanNotEquals,
	{"&&", string(object.BooleanKind), string(object.BooleanKind)}: booleanAnd,
	{"||", string(object.BooleanKind), string(object.BooleanKind)}: booleanOr,

	{"+", string(object.StringKind), string(object.StringKind)}:  stringAddition,
	{"+=", string(object.StringKind), string(object.StringKind)}: stringPlusEquals,
}

func evalInfixExpression(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)

	if object.IsError(left) {
		return left
	}

	right := Eval(node.Right, env)
	if object.IsError(right) {
		return right
	}

	operator := node.Operator

	if operator == string(token.ASSIGN) {
		return assignment(left, right, env)
	}

	operatorFnSignature := OperatorFnSignature{
		Operator: operator,
		LType:    object.UnwrapReferenceType(left.Type()).Signature(),
		RType:    object.UnwrapReferenceType(right.Type()).Signature(),
	}
	evalFn := infixEvalFns[operatorFnSignature]

	if evalFn == nil && object.IsNumber(left) && object.IsNumber(right) {
		return evalNumberInfixExpression(left, right, operator, env)
	}

	if evalFn == nil {
		return newError("operator %s not defined on types %s and %s", operatorFnSignature.Operator, operatorFnSignature.LType, operatorFnSignature.RType)
	}

	return evalFn(left, right, env)
}

func evalNumberInfixExpression(left object.Object, right object.Object, operator string, env *object.Environment) object.Object {
	leftNum := object.UnwrapReferenceObject(left).(*object.Number)
	rightNum := object.UnwrapReferenceObject(right).(*object.Number)
	switch operator {
	case string(token.PLUS):
		return numberAddition(leftNum, rightNum, env)
	case string(token.MINUS):
		return numberSubtraction(leftNum, rightNum, env)
	case string(token.ASTERISK):
		return numberMultiplication(leftNum, rightNum, env)
	case string(token.SLASH):
		return numberDivision(leftNum, rightNum, env)
	case string(token.B_SHIFT_L):
		return numberBitwiseShiftLeft(leftNum, rightNum, env)
	case string(token.B_SHIFT_R):
		return numberBitwiseShiftRight(leftNum, rightNum, env)
	case string(token.B_AND):
		return numberBitwiseAnd(leftNum, rightNum, env)
	case string(token.B_OR):
		return numberBitwiseOr(leftNum, rightNum, env)
	case string(token.B_XOR):
		return numberBitwiseXor(leftNum, rightNum, env)
	case string(token.EQ):
		return numberEquals(leftNum, rightNum, env)
	case string(token.NOT_EQ):
		return numberNotEquals(leftNum, rightNum, env)
	case string(token.GT):
		return numberGreaterThan(leftNum, rightNum, env)
	case string(token.LT):
		return numberLessThan(leftNum, rightNum, env)
	case string(token.LTE):
		return numberLessThanEqual(leftNum, rightNum, env)
	case string(token.GTE):
		return numberGreaterThanEqual(leftNum, rightNum, env)
	case string(token.PLUS_ASSIGN):
		return numberPlusEquals(left, rightNum, env)
	case string(token.MINUS_ASSIGN):
		return numberMinusEquals(left, rightNum, env)
	case string(token.ASTERISK_ASSIGN):
		return numberTimesEquals(left, rightNum, env)
	case string(token.SLASH_ASSIGN):
		return numberDivideEquals(left, rightNum, env)
	default:
		return newError("operator %s not defined on types %s and %s", operator, left.Type().Signature(), right.Type().Signature())
	}
}

func numberAddition(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)
	if err != nil {
		return newError(err.Error())
	}

	var sum *object.Number

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		sum = &object.Number{Value: object.Int64Bits(leftNum.GetInt64() + rightNum.GetInt64()), Kind: object.Int64Kind}
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		sum = &object.Number{Value: leftNum.GetUInt64() + rightNum.GetUInt64(), Kind: object.UInt64Kind}
	} else if object.IsFloat32(leftNum) {
		sum = &object.Number{Value: uint64(math.Float32bits(leftNum.GetFloat32() + rightNum.GetFloat32())), Kind: object.Float32Kind}
	} else if object.IsFloat64(leftNum) {
		sum = &object.Number{Value: math.Float64bits(leftNum.GetFloat64() + rightNum.GetFloat64()), Kind: object.Float64Kind}
	}

	castedSum, err := numberCast(sum, leftNum.Kind, EXPLICIT_CAST)

	if err != nil {
		return newError(err.Error())
	}

	return castedSum
}

func numberSubtraction(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)

	if err != nil {
		return newError(err.Error())
	}

	var difference *object.Number

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		difference = &object.Number{Value: object.Int64Bits(leftNum.GetInt64() - rightNum.GetInt64()), Kind: object.Int64Kind}
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		difference = &object.Number{Value: leftNum.GetUInt64() - rightNum.GetUInt64(), Kind: object.UInt64Kind}
	} else if object.IsFloat32(leftNum) {
		difference = &object.Number{Value: uint64(math.Float32bits(leftNum.GetFloat32() - rightNum.GetFloat32())), Kind: object.Float32Kind}
	} else if object.IsFloat64(leftNum) {
		difference = &object.Number{Value: math.Float64bits(leftNum.GetFloat64() - rightNum.GetFloat64()), Kind: object.Float64Kind}
	}

	castedDiffrence, err := numberCast(difference, leftNum.Kind, EXPLICIT_CAST)
	if err != nil {
		return newError(err.Error())
	}

	return castedDiffrence
}

func numberMultiplication(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)

	if err != nil {
		return newError(err.Error())
	}

	var product *object.Number

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		product = &object.Number{Value: object.Int64Bits(leftNum.GetInt64() * rightNum.GetInt64()), Kind: object.Int64Kind}
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		product = &object.Number{Value: leftNum.GetUInt64() * rightNum.GetUInt64(), Kind: object.UInt64Kind}
	} else if object.IsFloat32(leftNum) {
		product = &object.Number{Value: uint64(math.Float32bits(leftNum.GetFloat32() * rightNum.GetFloat32())), Kind: object.Float32Kind}
	} else if object.IsFloat64(leftNum) {
		product = &object.Number{Value: math.Float64bits(leftNum.GetFloat64() * rightNum.GetFloat64()), Kind: object.Float64Kind}
	}

	castedProduct, err := numberCast(product, leftNum.Kind, EXPLICIT_CAST)

	if err != nil {
		return newError(err.Error())
	}

	return castedProduct
}

func numberDivision(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)

	if err != nil {
		return newError(err.Error())
	}

	var quotient *object.Number

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		quotient = &object.Number{Value: object.Int64Bits(leftNum.GetInt64() / rightNum.GetInt64()), Kind: object.Int64Kind}
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		quotient = &object.Number{Value: leftNum.GetUInt64() / rightNum.GetUInt64(), Kind: object.UInt64Kind}
	} else if object.IsFloat32(leftNum) {
		quotient = &object.Number{Value: uint64(math.Float32bits(leftNum.GetFloat32() / rightNum.GetFloat32())), Kind: object.Float32Kind}
	} else if object.IsFloat64(leftNum) {
		quotient = &object.Number{Value: math.Float64bits(leftNum.GetFloat64() / rightNum.GetFloat64()), Kind: object.Float64Kind}
	}

	castedQuotient, err := numberCast(quotient, leftNum.Kind, EXPLICIT_CAST)

	if err != nil {
		return newError(err.Error())
	}

	return castedQuotient
}

func numberBitwiseShiftLeft(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	if object.IsFloat32(right) || object.IsFloat64(right) {
		return newError("Operator << not defined for %s and %s", left.Kind, right.Kind)
	} else if right.IsSigned() && right.GetInt64() < 0 {
		return newError("Operator << not defined on negative shift amount")
	}

	return &object.Number{Value: left.Value << right.GetUInt64(), Kind: left.Kind}
}

func numberBitwiseShiftRight(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	if object.IsFloat32(right) || object.IsFloat64(right) {
		return newError("Operator >> not defined for %s and %s", left.Kind, right.Kind)
	} else if right.IsSigned() && right.GetInt64() < 0 {
		return newError("Operator >> not defined on negative shift amount")
	}

	return &object.Number{Value: left.Value >> right.GetUInt64(), Kind: left.Kind}
}

func numberBitwiseAnd(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	return &object.Number{Value: left.Value & right.GetUInt64(), Kind: left.Kind}
}

func numberBitwiseOr(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	return &object.Number{Value: left.Value | right.GetUInt64(), Kind: left.Kind}
}

func numberBitwiseXor(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	return &object.Number{Value: left.Value ^ right.GetUInt64(), Kind: left.Kind}
}

func numberLessThan(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)

	if err != nil {
		return newError(err.Error())
	}

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		return nativeBoolToBooleanObject(leftNum.GetInt64() < rightNum.GetInt64())
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		return nativeBoolToBooleanObject(leftNum.GetUInt64() < rightNum.GetUInt64())
	} else if object.IsFloat32(leftNum) {
		return nativeBoolToBooleanObject(leftNum.GetFloat32() < rightNum.GetFloat32())
	} else if object.IsFloat64(leftNum) {
		return nativeBoolToBooleanObject(leftNum.GetFloat64() < rightNum.GetFloat64())
	}

	return NULL
}

func numberGreaterThan(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)

	if err != nil {
		return newError(err.Error())
	}

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		return nativeBoolToBooleanObject(leftNum.GetInt64() > rightNum.GetInt64())
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		return nativeBoolToBooleanObject(leftNum.GetUInt64() > rightNum.GetUInt64())
	} else if object.IsFloat32(leftNum) {
		return nativeBoolToBooleanObject(leftNum.GetFloat32() > rightNum.GetFloat32())
	} else if object.IsFloat64(leftNum) {
		return nativeBoolToBooleanObject(leftNum.GetFloat64() > rightNum.GetFloat64())
	}

	return NULL
}

func numberEquals(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	leftNum, rightNum, err := arithmeticCast(left, right)

	if err != nil {
		return newError(err.Error())
	}

	if object.IsInteger(leftNum) && leftNum.IsSigned() {
		return nativeBoolToBooleanObject(leftNum.GetInt64() == rightNum.GetInt64())
	} else if object.IsInteger(leftNum) && leftNum.IsUnsigned() {
		return nativeBoolToBooleanObject(leftNum.GetUInt64() == rightNum.GetUInt64())
	} else if object.IsFloat32(leftNum) {
		return nativeBoolToBooleanObject(leftNum.GetFloat32() == rightNum.GetFloat32())
	} else if object.IsFloat64(leftNum) {
		return nativeBoolToBooleanObject(leftNum.GetFloat64() == rightNum.GetFloat64())
	}

	return NULL
}

func numberLessThanEqual(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	return evalBangPrefixOperatorExpression(numberGreaterThan(left, right, env))
}

func numberGreaterThanEqual(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	return evalBangPrefixOperatorExpression(numberLessThan(left, right, env))
}

func numberNotEquals(left *object.Number, right *object.Number, env *object.Environment) object.Object {
	return evalBangPrefixOperatorExpression(numberEquals(left, right, env))
}

func numberPlusEquals(left object.Object, right *object.Number, env *object.Environment) object.Object {
	return assignment(left, numberAddition(object.UnwrapReferenceObject(left).(*object.Number), right, env), env)
}

func numberMinusEquals(left object.Object, right *object.Number, env *object.Environment) object.Object {
	return assignment(left, numberSubtraction(object.UnwrapReferenceObject(left).(*object.Number), right, env), env)
}

func numberTimesEquals(left object.Object, right *object.Number, env *object.Environment) object.Object {
	return assignment(left, numberMultiplication(object.UnwrapReferenceObject(left).(*object.Number), right, env), env)
}

func numberDivideEquals(left object.Object, right *object.Number, env *object.Environment) object.Object {
	return assignment(left, numberDivision(object.UnwrapReferenceObject(left).(*object.Number), right, env), env)
}

func assignment(left object.Object, right object.Object, env *object.Environment) object.Object {
	lvalue, ok := left.(object.ObjectReference)
	rvalue := object.UnwrapReferenceObject(right)
	lvalueType := object.UnwrapReferenceType(lvalue.GetValue().Type())
	rvalueType := object.UnwrapReferenceType(rvalue.Type())

	if !ok {
		return newError("Invalid lvalue %s", left.Inspect())
	}
	if lvalue.Type().IsConstant() {
		return newError("Cannot assign to const variable")
	}

	if lvalueType.Signature() != rvalueType.Signature() {
		cast := typeCast(rvalue, lvalueType, EXPLICIT_CAST)
		if !object.IsError(cast) {
			rvalue = cast
		} else {
			return cast
		}
	}

	_, err := lvalue.UpdateValue(rvalue)
	if err != nil {
		return newError(err.Error())
	}
	return lvalue
}

func booleanEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftBool := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value == rightBool.Value)
}

func booleanNotEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftBool := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value != rightBool.Value)
}

func booleanAnd(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftBool := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value && rightBool.Value)
}

func booleanOr(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftBool := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value || rightBool.Value)
}

func stringAddition(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftStr := object.UnwrapReferenceObject(left).(*object.String)
	rightStr := object.UnwrapReferenceObject(right).(*object.String)

	return &object.String{Value: leftStr.Value + rightStr.Value}
}

func stringPlusEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, stringAddition(left, right, env), env)
}
