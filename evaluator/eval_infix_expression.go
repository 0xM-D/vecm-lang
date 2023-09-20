package evaluator

import (
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

	{"&", string(object.Int64Kind), string(object.Int64Kind)}:  integerBitwiseAnd,
	{"|", string(object.Int64Kind), string(object.Int64Kind)}:  integerBitwiseOr,
	{"^", string(object.Int64Kind), string(object.Int64Kind)}:  integerBitwiseXor,
	{">>", string(object.Int64Kind), string(object.Int64Kind)}: integerBitwiseShiftRight,
	{"<<", string(object.Int64Kind), string(object.Int64Kind)}: integerBitwiseShiftLeft,

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
	switch operator {
	case string(token.PLUS):
		return numberAddition(left, right, env)
	case string(token.MINUS):
		return numberSubtraction(left, right, env)
	case string(token.ASTERISK):
		return numberMultiplication(left, right, env)
	case string(token.SLASH):
		return numberDivision(left, right, env)
	case string(token.EQ):
		return numberEquals(left, right, env)
	case string(token.NOT_EQ):
		return numberNotEquals(left, right, env)
	case string(token.GT):
		return numberGreaterThan(left, right, env)
	case string(token.LT):
		return numberLessThan(left, right, env)
	case string(token.LTE):
		return numberLessThanEqual(left, right, env)
	case string(token.GTE):
		return numberGreaterThanEqual(left, right, env)
	case string(token.PLUS_ASSIGN):
		return numberPlusEquals(left, right, env)
	case string(token.MINUS_ASSIGN):
		return numberMinusEquals(left, right, env)
	case string(token.ASTERISK_ASSIGN):
		return numberTimesEquals(left, right, env)
	case string(token.SLASH_ASSIGN):
		return numberDivideEquals(left, right, env)

	default:
		return newError("operator %s not defined on types %s and %s", operator, left.Type().Signature(), right.Type().Signature())
	}
}

func add[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Number[T] {
	return &object.Number[T]{Value: a.(*object.Number[T]).Value + b.(*object.Number[T]).Value}
}

func subtract[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Number[T] {
	return &object.Number[T]{Value: a.(*object.Number[T]).Value - b.(*object.Number[T]).Value}
}

func multiply[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Number[T] {
	return &object.Number[T]{Value: a.(*object.Number[T]).Value * b.(*object.Number[T]).Value}
}

func divide[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Number[T] {
	return &object.Number[T]{Value: a.(*object.Number[T]).Value / b.(*object.Number[T]).Value}
}

func lessThan[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(*object.Number[T]).Value < b.(*object.Number[T]).Value)
}

func greaterThan[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(*object.Number[T]).Value > b.(*object.Number[T]).Value)
}

func equals[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(*object.Number[T]).Value == b.(*object.Number[T]).Value)
}

func integerBitwiseAnd(a object.Object, b object.Object, env *object.Environment) object.Object {
	return &object.Number[int64]{Value: a.(*object.Number[int64]).Value & b.(*object.Number[int64]).Value}
}

func integerBitwiseOr(a object.Object, b object.Object, env *object.Environment) object.Object {
	return &object.Number[int64]{Value: a.(*object.Number[int64]).Value | b.(*object.Number[int64]).Value}
}

func integerBitwiseXor(a object.Object, b object.Object, env *object.Environment) object.Object {
	return &object.Number[int64]{Value: a.(*object.Number[int64]).Value ^ b.(*object.Number[int64]).Value}
}

func integerBitwiseShiftLeft(a object.Object, b object.Object, env *object.Environment) object.Object {
	return &object.Number[int64]{Value: a.(*object.Number[int64]).Value << b.(*object.Number[int64]).Value}
}

func integerBitwiseShiftRight(a object.Object, b object.Object, env *object.Environment) object.Object {
	return &object.Number[int64]{Value: a.(*object.Number[int64]).Value >> b.(*object.Number[int64]).Value}
}

func numberAddition(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return add[int8](leftNum, rightNum)
	case object.Int16Kind:
		return add[int16](leftNum, rightNum)
	case object.Int32Kind:
		return add[int32](leftNum, rightNum)
	case object.Int64Kind:
		return add[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return add[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return add[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return add[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return add[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return add[float32](leftNum, rightNum)
	case object.Float64Kind:
		return add[float64](leftNum, rightNum)
	}

	return nil
}

func numberSubtraction(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return subtract[int8](leftNum, rightNum)
	case object.Int16Kind:
		return subtract[int16](leftNum, rightNum)
	case object.Int32Kind:
		return subtract[int32](leftNum, rightNum)
	case object.Int64Kind:
		return subtract[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return subtract[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return subtract[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return subtract[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return subtract[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return subtract[float32](leftNum, rightNum)
	case object.Float64Kind:
		return subtract[float64](leftNum, rightNum)
	}

	return nil
}

func numberMultiplication(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return multiply[int8](leftNum, rightNum)
	case object.Int16Kind:
		return multiply[int16](leftNum, rightNum)
	case object.Int32Kind:
		return multiply[int32](leftNum, rightNum)
	case object.Int64Kind:
		return multiply[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return multiply[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return multiply[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return multiply[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return multiply[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return multiply[float32](leftNum, rightNum)
	case object.Float64Kind:
		return multiply[float64](leftNum, rightNum)
	}

	return nil
}

func numberDivision(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return divide[int8](leftNum, rightNum)
	case object.Int16Kind:
		return divide[int16](leftNum, rightNum)
	case object.Int32Kind:
		return divide[int32](leftNum, rightNum)
	case object.Int64Kind:
		return divide[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return divide[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return divide[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return divide[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return divide[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return divide[float32](leftNum, rightNum)
	case object.Float64Kind:
		return divide[float64](leftNum, rightNum)
	}

	return nil
}

func numberLessThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return lessThan[int8](leftNum, rightNum)
	case object.Int16Kind:
		return lessThan[int16](leftNum, rightNum)
	case object.Int32Kind:
		return lessThan[int32](leftNum, rightNum)
	case object.Int64Kind:
		return lessThan[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return lessThan[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return lessThan[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return lessThan[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return lessThan[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return lessThan[float32](leftNum, rightNum)
	case object.Float64Kind:
		return lessThan[float64](leftNum, rightNum)
	}

	return nil
}

func numberLessThanEqual(left object.Object, right object.Object, env *object.Environment) object.Object {
	return evalBangOperatorExpression(numberGreaterThan(left, right, env))
}

func numberGreaterThanEqual(left object.Object, right object.Object, env *object.Environment) object.Object {
	return evalBangOperatorExpression(numberLessThan(left, right, env))
}

func numberGreaterThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return greaterThan[int8](leftNum, rightNum)
	case object.Int16Kind:
		return greaterThan[int16](leftNum, rightNum)
	case object.Int32Kind:
		return greaterThan[int32](leftNum, rightNum)
	case object.Int64Kind:
		return greaterThan[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return greaterThan[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return greaterThan[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return greaterThan[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return greaterThan[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return greaterThan[float32](leftNum, rightNum)
	case object.Float64Kind:
		return greaterThan[float64](leftNum, rightNum)
	}

	return nil
}

func numberEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	nums := castToLargerNumberType(left, right)
	leftNum := nums[0]
	rightNum := nums[1]

	switch leftNum.Type().Kind() {
	case object.Int8Kind:
		return equals[int8](leftNum, rightNum)
	case object.Int16Kind:
		return equals[int16](leftNum, rightNum)
	case object.Int32Kind:
		return equals[int32](leftNum, rightNum)
	case object.Int64Kind:
		return equals[int64](leftNum, rightNum)
	case object.UInt8Kind:
		return equals[uint8](leftNum, rightNum)
	case object.UInt16Kind:
		return equals[uint16](leftNum, rightNum)
	case object.UInt32Kind:
		return equals[uint32](leftNum, rightNum)
	case object.UInt64Kind:
		return equals[uint64](leftNum, rightNum)
	case object.Float32Kind:
		return equals[float32](leftNum, rightNum)
	case object.Float64Kind:
		return equals[float64](leftNum, rightNum)
	}

	return nil
}

func numberNotEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return evalBangOperatorExpression(numberEquals(left, right, env))
}

func numberPlusEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, numberAddition(left, right, env), env)
}

func numberMinusEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, numberSubtraction(left, right, env), env)
}

func numberTimesEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, numberMultiplication(left, right, env), env)
}

func numberDivideEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, numberDivision(left, right, env), env)
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
		cast := typeCast(rvalue, lvalueType, IMPLICIT_CAST)
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
