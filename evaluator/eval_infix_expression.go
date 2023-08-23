package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

type OperatorFnSignature struct {
	Operator string
	LType    string
	RType    string
}

type InfixEvalFn func(object.Object, object.Object, *object.Environment) object.Object

var infixEvalFns = map[OperatorFnSignature]InfixEvalFn{
	{"=", "int", "int"}:       assignment,
	{"=", "string", "string"}: assignment,
	{"=", "bool", "bool"}:     assignment,

	{"+", "int", "int"}:  integerAddition,
	{"-", "int", "int"}:  integerSubtraction,
	{"*", "int", "int"}:  integerMultiplication,
	{"/", "int", "int"}:  integerDivision,
	{"==", "int", "int"}: integerEquals,
	{"!=", "int", "int"}: integerNotEquals,
	{"<", "int", "int"}:  integerLessThan,
	{">", "int", "int"}:  integerGreaterThan,
	{"+=", "int", "int"}: integerPlusEquals,
	{"-=", "int", "int"}: integerMinusEquals,
	{"*=", "int", "int"}: integerTimesEquals,
	{"/=", "int", "int"}: integerDivideEquals,
	{"&", "int", "int"}:  integerDivideEquals,
	{"|", "int", "int"}:  integerDivideEquals,
	{"^", "int", "int"}:  integerDivideEquals,
	{"<<", "int", "int"}: integerDivideEquals,
	{">>", "int", "int"}: integerDivideEquals,

	{"==", "bool", "bool"}: booleanEquals,
	{"!=", "bool", "bool"}: booleanNotEquals,
	{"&&", "bool", "bool"}: booleanEquals,
	{"||", "bool", "bool"}: booleanNotEquals,

	{"+", "string", "string"}:  stringAddition,
	{"+=", "string", "string"}: stringPlusEquals,
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
	operatorFnSignature := OperatorFnSignature{Operator: operator, LType: left.Type().Signature(), RType: right.Type().Signature()}
	evalFn := infixEvalFns[operatorFnSignature]
	if evalFn == nil {
		return newError("operator %s not defined on types %s and %s", operatorFnSignature.Operator, operatorFnSignature.LType, operatorFnSignature.RType)
	}

	return evalFn(left, right, env)
}

func integerAddition(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return &object.Integer{Value: leftInt.Value + rightInt.Value}
}

func integerSubtraction(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return &object.Integer{Value: leftInt.Value - rightInt.Value}
}

func integerDivision(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return &object.Integer{Value: leftInt.Value / rightInt.Value}
}

func integerMultiplication(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return &object.Integer{Value: leftInt.Value * rightInt.Value}
}

func integerLessThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value < rightInt.Value)
}

func integerGreaterThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value > rightInt.Value)
}

func integerEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value == rightInt.Value)
}

func integerNotEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := object.UnwrapReferenceObject(left).(*object.Integer)
	rightInt := object.UnwrapReferenceObject(right).(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value != rightInt.Value)
}

func integerPlusEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, integerAddition(left, right, env), env)
}

func integerMinusEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, integerSubtraction(left, right, env), env)
}

func integerTimesEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, integerMultiplication(left, right, env), env)
}

func integerDivideEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	return assignment(left, integerDivision(left, right, env), env)
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
