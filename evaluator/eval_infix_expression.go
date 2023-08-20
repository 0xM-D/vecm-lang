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

	if object.IsNumber(left) && object.IsNumber(right) {
		return evalNumberInfixExpression(left, right, operator, env)
	}

	operatorFnSignature := OperatorFnSignature{Operator: operator, LType: left.Type().Signature(), RType: right.Type().Signature()}
	evalFn := infixEvalFns[operatorFnSignature]
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

func add[T int64 | float32 | float64](a object.Object, b object.Object) object.Number[T] {
	return object.Number[T]{Value: a.(object.Number[T]).Value + b.(object.Number[T]).Value}
}

func subtract[T int64 | float32 | float64](a object.Object, b object.Object) object.Number[T] {
	return object.Number[T]{Value: a.(object.Number[T]).Value - b.(object.Number[T]).Value}
}

func multiply[T int64 | float32 | float64](a object.Object, b object.Object) object.Number[T] {
	return object.Number[T]{Value: a.(object.Number[T]).Value * b.(object.Number[T]).Value}
}

func divide[T int64 | float32 | float64](a object.Object, b object.Object) object.Number[T] {
	return object.Number[T]{Value: a.(object.Number[T]).Value / b.(object.Number[T]).Value}
}

func lessThan[T int64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(object.Number[T]).Value < b.(object.Number[T]).Value)
}

func greaterThan[T int64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(object.Number[T]).Value > b.(object.Number[T]).Value)
}

func equals[T int64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(object.Number[T]).Value == b.(object.Number[T]).Value)
}

func notEquals[T int64 | float32 | float64](a object.Object, b object.Object) *object.Boolean {
	return nativeBoolToBooleanObject(a.(object.Number[T]).Value != b.(object.Number[T]).Value)
}

func numberAddition(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return add[int64](leftNum, rightNum)
	case object.Float32Kind:
		return add[float32](leftNum, rightNum)
	case object.Float64Kind:
		return add[float64](leftNum, rightNum)
	}

	return nil
}

func numberSubtraction(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return subtract[int64](leftNum, rightNum)
	case object.Float32Kind:
		return subtract[float32](leftNum, rightNum)
	case object.Float64Kind:
		return subtract[float64](leftNum, rightNum)
	}

	return nil
}

func numberMultiplication(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return multiply[int64](leftNum, rightNum)
	case object.Float32Kind:
		return multiply[float32](leftNum, rightNum)
	case object.Float64Kind:
		return multiply[float64](leftNum, rightNum)
	}

	return nil
}

func numberDivision(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return divide[int64](leftNum, rightNum)
	case object.Float32Kind:
		return divide[float32](leftNum, rightNum)
	case object.Float64Kind:
		return divide[float64](leftNum, rightNum)
	}

	return nil
}

func numberLessThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return lessThan[int64](leftNum, rightNum)
	case object.Float32Kind:
		return lessThan[float32](leftNum, rightNum)
	case object.Float64Kind:
		return lessThan[float64](leftNum, rightNum)
	}

	return nil
}

func numberGreaterThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return greaterThan[int64](leftNum, rightNum)
	case object.Float32Kind:
		return greaterThan[float32](leftNum, rightNum)
	case object.Float64Kind:
		return greaterThan[float64](leftNum, rightNum)
	}

	return nil
}

func numberEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return equals[int64](leftNum, rightNum)
	case object.Float32Kind:
		return equals[float32](leftNum, rightNum)
	case object.Float64Kind:
		return equals[float64](leftNum, rightNum)
	}

	return nil
}

func numberNotEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftNum, rightNum, kind := castToLargerNumberType(left, right)

	switch kind {
	case object.IntegerKind:
		return notEquals[int64](leftNum, rightNum)
	case object.Float32Kind:
		return notEquals[float32](leftNum, rightNum)
	case object.Float64Kind:
		return notEquals[float64](leftNum, rightNum)
	}

	return nil
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
	lvalue, ok := left.(*object.ObjectReference)
	if !ok {
		return newError("Invalid lvalue %s", left.Inspect())
	}

	newReference := env.Set(lvalue.Identifier, object.UnwrapReferenceObject(right))
	if newReference == nil {
		return newError("Cannot assign to const variable")
	}

	return newReference
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

func stringAddition(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftStr := object.UnwrapReferenceObject(left).(*object.String)
	rightStr := object.UnwrapReferenceObject(right).(*object.String)

	return &object.String{Value: leftStr.Value + rightStr.Value}
}

func stringPlusEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, stringAddition(left, right, env), env)
}
