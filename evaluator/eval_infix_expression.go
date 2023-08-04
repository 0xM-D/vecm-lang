package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

type OperatorFnSignature struct {
	Operator string
	LType    object.ObjectType
	RType    object.ObjectType
}

type InfixEvalFn func(object.Object, object.Object, *object.Environment) object.Object

var infixEvalFns = map[OperatorFnSignature]InfixEvalFn{
	{"+", object.INTEGER_OBJ(), object.INTEGER_OBJ()}:  integerAddition,
	{"-", object.INTEGER_OBJ(), object.INTEGER_OBJ()}:  integerSubtraction,
	{"*", object.INTEGER_OBJ(), object.INTEGER_OBJ()}:  integerMultiplication,
	{"/", object.INTEGER_OBJ(), object.INTEGER_OBJ()}:  integerDivision,
	{"==", object.INTEGER_OBJ(), object.INTEGER_OBJ()}: integerEquals,
	{"!=", object.INTEGER_OBJ(), object.INTEGER_OBJ()}: integerNotEquals,
	{"<", object.INTEGER_OBJ(), object.INTEGER_OBJ()}:  integerLessThan,
	{">", object.INTEGER_OBJ(), object.INTEGER_OBJ()}:  integerGreaterThan,
}

func evalInfixExpression2(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if object.IsError(left) {
		return left
	}

	right := Eval(node.Right, env)
	if object.IsError(right) {
		return right
	}
	operator := node.Operator

	operatorFnSignature := OperatorFnSignature{Operator: operator, LType: left.Type(), RType: right.Type()}
	evalFn := infixEvalFns[operatorFnSignature]
	if evalFn == nil {
		return newError("operator %s not defined on types %s and %s", operatorFnSignature.Operator, operatorFnSignature.LType, operatorFnSignature.RType)
	}

	return evalFn(left, right, env)
}

func integerAddition(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return &object.Integer{Value: leftInt.Value + rightInt.Value}
}

func integerSubtraction(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return &object.Integer{Value: leftInt.Value - rightInt.Value}
}

func integerDivision(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return &object.Integer{Value: leftInt.Value / rightInt.Value}
}

func integerMultiplication(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return &object.Integer{Value: leftInt.Value * rightInt.Value}
}

func integerLessThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value < rightInt.Value)
}

func integerGreaterThan(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value > rightInt.Value)
}

func integerEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value == rightInt.Value)
}

func integerNotEquals(left object.Object, right object.Object, env *object.Environment) object.Object {
	leftInt := left.(*object.Integer)
	rightInt := right.(*object.Integer)

	return nativeBoolToBooleanObject(leftInt.Value != rightInt.Value)
}
