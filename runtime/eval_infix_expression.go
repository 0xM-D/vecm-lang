package runtime

import (
	"fmt"
	"math"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
	"github.com/DustTheory/interpreter/token"
	"github.com/pkg/errors"
)

type InfixEvalSettings struct {
	InfixEvalFns map[OperatorFnSignature]InfixEvalFn
}

var infixEvalSettings = newInfixEvalSettings()

func newInfixEvalSettings() *InfixEvalSettings {
	return &InfixEvalSettings{
		InfixEvalFns: map[OperatorFnSignature]InfixEvalFn{
			{"==", string(object.BooleanKind), string(object.BooleanKind)}: booleanEquals,
			{"!=", string(object.BooleanKind), string(object.BooleanKind)}: booleanNotEquals,
			{"&&", string(object.BooleanKind), string(object.BooleanKind)}: booleanAnd,
			{"||", string(object.BooleanKind), string(object.BooleanKind)}: booleanOr,

			{"+", string(object.StringKind), string(object.StringKind)}:  stringAddition,
			{"+=", string(object.StringKind), string(object.StringKind)}: stringPlusEquals,
		},
	}
}

type OperatorFnSignature struct {
	Operator string
	LType    string
	RType    string
}

type InfixEvalFn func(object.Object, object.Object, *object.Environment) (object.Object, error)

func (r *Runtime) evalInfixExpression(node *ast.InfixExpression, env *object.Environment) (object.Object, error) {
	left, err := r.Eval(node.Left, env)

	if err != nil {
		return nil, err
	}

	right, err := r.Eval(node.Right, env)
	if err != nil {
		return nil, err
	}

	operator := node.Operator

	if operator == string(token.Assign) {
		return assignment(left, right, env)
	}

	operatorFnSignature := OperatorFnSignature{
		Operator: operator,
		LType:    object.UnwrapReferenceType(left.Type()).Signature(),
		RType:    object.UnwrapReferenceType(right.Type()).Signature(),
	}
	evalFn := infixEvalSettings.InfixEvalFns[operatorFnSignature]

	if evalFn == nil && object.IsNumber(left) && object.IsNumber(right) {
		return r.evalNumberInfixExpression(left, right, operator, env)
	}

	if evalFn == nil {
		return nil, fmt.Errorf("operator %s not defined on types %s and %s",
			operatorFnSignature.Operator,
			operatorFnSignature.LType,
			operatorFnSignature.RType)
	}

	return evalFn(left, right, env)
}

func (r *Runtime) evalNumberInfixExpression(
	left object.Object,
	right object.Object,
	operator string,
	env *object.Environment,
) (object.Object, error) {
	leftNum, ok := object.UnwrapReferenceObject(left).(*object.Number)
	if !ok {
		return nil, fmt.Errorf("invalid left operand type: expected *object.Number, got %T", left)
	}
	rightNum, ok := object.UnwrapReferenceObject(right).(*object.Number)
	if !ok {
		return nil, fmt.Errorf("invalid right operand type: expected *object.Number, got %T", left)
	}
	switch operator {
	case string(token.Plus):
		return numberAddition(leftNum, rightNum)
	case string(token.Minus):
		return numberSubtraction(leftNum, rightNum, env)
	case string(token.Asterisk):
		return numberMultiplication(leftNum, rightNum, env)
	case string(token.Slash):
		return numberDivision(leftNum, rightNum, env)
	case string(token.BitwiseShiftL):
		return numberBitwiseShiftLeft(leftNum, rightNum, env)
	case string(token.BitwiseShiftR):
		return numberBitwiseShiftRight(leftNum, rightNum, env)
	case string(token.BitwiseAnd):
		return numberBitwiseAnd(leftNum, rightNum, env), nil
	case string(token.BitwiseOr):
		return numberBitwiseOr(leftNum, rightNum, env), nil
	case string(token.BitwiseXor):
		return numberBitwiseXor(leftNum, rightNum, env), nil
	case string(token.Eq):
		return numberEquals(leftNum, rightNum, env)
	case string(token.NotEq):
		return numberNotEquals(leftNum, rightNum, env)
	case string(token.Gt):
		return numberGreaterThan(leftNum, rightNum, env)
	case string(token.Lt):
		return numberLessThan(leftNum, rightNum, env)
	case string(token.Lte):
		return numberLessThanEqual(leftNum, rightNum, env)
	case string(token.Gte):
		return numberGreaterThanEqual(leftNum, rightNum, env)
	case string(token.PlusAssign):
		return numberPlusEquals(left, rightNum, env)
	case string(token.MinusAssign):
		return numberMinusEquals(left, rightNum, env)
	case string(token.AsteriskAssing):
		return numberTimesEquals(left, rightNum, env)
	case string(token.SlashAssign):
		return numberDivideEquals(left, rightNum, env)
	default:
		return nil, fmt.Errorf("operator %s not defined on types %s and %s",
			operator,
			left.Type().Signature(),
			right.Type().Signature())
	}
}

func numberAddition(left *object.Number, right *object.Number) (object.Object, error) {
	return performArithmeticOperation(
		left, right,
		func(a, b int64) int64 { return a + b },
		func(a, b uint64) uint64 { return a + b },
		func(a, b float32) float32 { return a + b },
		func(a, b float64) float64 { return a + b },
	)
}

func numberSubtraction(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	return performArithmeticOperation(
		left, right,
		func(a, b int64) int64 { return a - b },
		func(a, b uint64) uint64 { return a - b },
		func(a, b float32) float32 { return a - b },
		func(a, b float64) float64 { return a - b },
	)
}

func numberMultiplication(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	return performArithmeticOperation(
		left, right,
		func(a, b int64) int64 { return a * b },
		func(a, b uint64) uint64 { return a * b },
		func(a, b float32) float32 { return a * b },
		func(a, b float64) float64 { return a * b },
	)
}

func numberDivision(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	return performArithmeticOperation(
		left, right,
		func(a, b int64) int64 { return a / b },
		func(a, b uint64) uint64 { return a / b },
		func(a, b float32) float32 { return a / b },
		func(a, b float64) float64 { return a / b },
	)
}

func numberBitwiseShiftLeft(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	if object.IsFloat32(right) || object.IsFloat64(right) {
		return nil, fmt.Errorf("operator << not defined for %s and %s", left.Kind, right.Kind)
	} else if right.IsSigned() && right.GetInt64() < 0 {
		return nil, errors.New("operator << not defined on negative shift amount")
	}

	return &object.Number{Value: left.Value << right.GetUInt64(), Kind: left.Kind}, nil
}

func numberBitwiseShiftRight(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	if object.IsFloat32(right) || object.IsFloat64(right) {
		return nil, fmt.Errorf("operator >> not defined for %s and %s", left.Kind, right.Kind)
	} else if right.IsSigned() && right.GetInt64() < 0 {
		return nil, errors.New("operator >> not defined on negative shift amount")
	}

	return &object.Number{Value: left.Value >> right.GetUInt64(), Kind: left.Kind}, nil
}

func numberBitwiseAnd(left *object.Number, right *object.Number, _ *object.Environment) object.Object {
	return &object.Number{Value: left.Value & right.GetUInt64(), Kind: left.Kind}
}

func numberBitwiseOr(left *object.Number, right *object.Number, _ *object.Environment) object.Object {
	return &object.Number{Value: left.Value | right.GetUInt64(), Kind: left.Kind}
}

func numberBitwiseXor(left *object.Number, right *object.Number, _ *object.Environment) object.Object {
	return &object.Number{Value: left.Value ^ right.GetUInt64(), Kind: left.Kind}
}

func numberLessThan(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	return performComparisonOperation(
		left, right,
		func(a, b int64) bool { return a < b },
		func(a, b uint64) bool { return a < b },
		func(a, b float32) bool { return a < b },
		func(a, b float64) bool { return a < b },
	)
}

func numberGreaterThan(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	return performComparisonOperation(
		left, right,
		func(a, b int64) bool { return a > b },
		func(a, b uint64) bool { return a > b },
		func(a, b float32) bool { return a > b },
		func(a, b float64) bool { return a > b },
	)
}

func numberEquals(left *object.Number, right *object.Number, _ *object.Environment) (object.Object, error) {
	return performComparisonOperation(
		left, right,
		func(a, b int64) bool { return a == b },
		func(a, b uint64) bool { return a == b },
		func(a, b float32) bool { return a == b },
		func(a, b float64) bool { return a == b },
	)
}

func numberLessThanEqual(left *object.Number, right *object.Number, env *object.Environment) (object.Object, error) {
	greaterThan, err := numberGreaterThan(left, right, env)
	if err != nil {
		return nil, err
	}
	return evalBangPrefixOperatorExpression(greaterThan)
}

func numberGreaterThanEqual(left *object.Number, right *object.Number, env *object.Environment) (object.Object, error) {
	lessThan, err := numberLessThan(left, right, env)
	if err != nil {
		return nil, err
	}
	return evalBangPrefixOperatorExpression(lessThan)
}

func numberNotEquals(left *object.Number, right *object.Number, env *object.Environment) (object.Object, error) {
	equals, err := numberEquals(left, right, env)
	if err != nil {
		return nil, err
	}
	return evalBangPrefixOperatorExpression(equals)
}

func numberPlusEquals(left object.Object, right *object.Number, env *object.Environment) (object.Object, error) {
	sum, err := numberAddition(object.UnwrapReferenceObject(left).(*object.Number), right)
	if err != nil {
		return nil, err
	}
	return assignment(left, sum, env)
}

func numberMinusEquals(left object.Object, right *object.Number, env *object.Environment) (object.Object, error) {
	difference, err := numberSubtraction(object.UnwrapReferenceObject(left).(*object.Number), right, env)
	if err != nil {
		return nil, err
	}
	return assignment(left, difference, env)
}

func numberTimesEquals(left object.Object, right *object.Number, env *object.Environment) (object.Object, error) {
	product, err := numberMultiplication(object.UnwrapReferenceObject(left).(*object.Number), right, env)
	if err != nil {
		return nil, err
	}
	return assignment(left, product, env)
}

func numberDivideEquals(left object.Object, right *object.Number, env *object.Environment) (object.Object, error) {
	quotient, err := numberDivision(object.UnwrapReferenceObject(left).(*object.Number), right, env)
	if err != nil {
		return nil, err
	}
	return assignment(left, quotient, env)
}

func assignment(left object.Object, right object.Object, _ *object.Environment) (object.Object, error) {
	lvalue, ok := left.(object.Reference)
	rvalue := object.UnwrapReferenceObject(right)
	lvalueType := object.UnwrapReferenceType(lvalue.GetValue().Type())
	rvalueType := object.UnwrapReferenceType(rvalue.Type())

	if !ok {
		return nil, fmt.Errorf("invalid lvalue %s", left.Inspect())
	}
	if lvalue.Type().IsConstant() {
		return nil, errors.New("cannot assign to const variable")
	}

	if lvalueType.Signature() != rvalueType.Signature() {
		cast, err := typeCast(rvalue, lvalueType, ExplicitCast)
		if err != nil {
			return nil, err
		}
		rvalue = cast
	}

	_, err := lvalue.UpdateValue(rvalue)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update lvalue")
	}

	return lvalue, nil
}

func booleanEquals(left object.Object, right object.Object, _ *object.Environment) (object.Object, error) {
	leftBool, _ := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool, _ := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value == rightBool.Value), nil
}

func booleanNotEquals(left object.Object, right object.Object, _ *object.Environment) (object.Object, error) {
	leftBool, _ := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool, _ := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value != rightBool.Value), nil
}

func booleanAnd(left object.Object, right object.Object, _ *object.Environment) (object.Object, error) {
	leftBool, _ := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool, _ := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value && rightBool.Value), nil
}

func booleanOr(left object.Object, right object.Object, _ *object.Environment) (object.Object, error) {
	leftBool, _ := object.UnwrapReferenceObject(left).(*object.Boolean)
	rightBool, _ := object.UnwrapReferenceObject(right).(*object.Boolean)

	return nativeBoolToBooleanObject(leftBool.Value || rightBool.Value), nil
}

func stringAddition(left object.Object, right object.Object, _ *object.Environment) (object.Object, error) {
	leftStr, _ := object.UnwrapReferenceObject(left).(*object.String)
	rightStr, _ := object.UnwrapReferenceObject(right).(*object.String)

	return &object.String{Value: leftStr.Value + rightStr.Value}, nil
}

func stringPlusEquals(left object.Object, right object.Object, env *object.Environment) (object.Object, error) {
	added, err := stringAddition(left, right, env)
	if err != nil {
		return nil, err
	}
	return assignment(left, added, env)
}

func performArithmeticOperation(
	left *object.Number,
	right *object.Number,
	operation func(int64, int64) int64,
	unsignedOperation func(uint64, uint64) uint64,
	float32Operation func(float32, float32) float32,
	float64Operation func(float64, float64) float64,
) (*object.Number, error) {
	leftNum, rightNum, err := arithmeticCast(left, right)
	if err != nil {
		return nil, err
	}

	var result *object.Number

	switch {
	case object.IsInteger(leftNum) && leftNum.IsSigned():
		result = &object.Number{
			Value: object.Int64Bits(operation(leftNum.GetInt64(), rightNum.GetInt64())),
			Kind:  object.Int64Kind,
		}
	case object.IsInteger(leftNum) && leftNum.IsUnsigned():
		result = &object.Number{
			Value: unsignedOperation(leftNum.GetUInt64(), rightNum.GetUInt64()),
			Kind:  object.UInt64Kind,
		}
	case object.IsFloat32(leftNum):
		result = &object.Number{
			Value: uint64(math.Float32bits(float32Operation(leftNum.GetFloat32(), rightNum.GetFloat32()))),
			Kind:  object.Float32Kind,
		}
	case object.IsFloat64(leftNum):
		result = &object.Number{
			Value: math.Float64bits(float64Operation(leftNum.GetFloat64(), rightNum.GetFloat64())),
			Kind:  object.Float64Kind,
		}
	}

	castedResult, err := numberCast(result, leftNum.Kind, ExplicitCast)
	if err != nil {
		return nil, err
	}

	return castedResult, nil
}

func performComparisonOperation(
	left *object.Number,
	right *object.Number,
	operation func(int64, int64) bool,
	unsignedOperation func(uint64, uint64) bool,
	float32Operation func(float32, float32) bool,
	float64Operation func(float64, float64) bool,
) (object.Object, error) {
	leftNum, rightNum, err := arithmeticCast(left, right)
	if err != nil {
		return nil, err
	}

	var result bool

	switch {
	case object.IsInteger(leftNum) && leftNum.IsSigned():
		result = operation(leftNum.GetInt64(), rightNum.GetInt64())
	case object.IsInteger(leftNum) && leftNum.IsUnsigned():
		result = unsignedOperation(leftNum.GetUInt64(), rightNum.GetUInt64())
	case object.IsFloat32(leftNum):
		result = float32Operation(leftNum.GetFloat32(), rightNum.GetFloat32())
	case object.IsFloat64(leftNum):
		result = float64Operation(leftNum.GetFloat64(), rightNum.GetFloat64())
	}

	return nativeBoolToBooleanObject(result), nil
}
