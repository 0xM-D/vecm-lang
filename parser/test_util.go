package parser

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/0xM-D/interpreter/ast"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value *big.Int) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value.Cmp(value) != 0 {
		t.Errorf("integ.Value not %d. got=%s", value, integ.Value.String())
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

func testFloat32Literal(t *testing.T, fl ast.Expression, value float32) bool {
	float, ok := fl.(*ast.Float32Literal)
	if !ok {
		t.Errorf("fl not *ast.Float32Literal. got=%T", fl)
		return false
	}

	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}

	return true
}

func testFloat64Literal(t *testing.T, fl ast.Expression, value float64) bool {
	float, ok := fl.(*ast.Float64Literal)
	if !ok {
		t.Errorf("fl not *ast.Float64Literal. got=%T", fl)
		return false
	}

	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func testStringLiteral(t *testing.T, exp ast.Expression, value string) bool {
	str, ok := exp.(*ast.StringLiteral)
	if !ok {
		t.Errorf("str not *ast.StringLiteral. got=%T", exp)
		return false
	}

	if str.Value != value {
		t.Errorf("str.Value not %s. got=%s", value, str.Value)
		return false
	}

	if str.TokenLiteral() != value {
		t.Errorf("str.TokenLiteral not %s. got=%s", value, str.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	boolean, ok := exp.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("be not *ast.Boolean. got=%T", exp)
		return false
	}

	if boolean.Value != value {
		t.Errorf("integ.Value not %t. got=%t", value, boolean.Value)
		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("integ.TokenLiteral not %t. got=%s", value, boolean.TokenLiteral())
		return false
	}

	return true
}

type TestIdentifier struct {
	string
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {

	expectedKind := reflect.TypeOf(expected).Kind()
	if expectedKind == reflect.Array || expectedKind == reflect.Slice {

		arrayLiteral := exp.(*ast.NewExpression)
		expectedElements := expected.([]interface{})

		for i, exp := range arrayLiteral.InitializationList {
			if !testLiteralExpression(t, exp, expectedElements[i]) {
				return false
			}
		}

		return true
	}

	switch v := expected.(type) {
	case *big.Int:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testStringLiteral(t, exp, v)
	case TestIdentifier:
		return testIdentifier(t, exp, v.string)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%T %T", exp, expected)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func testAccessExpression(t *testing.T, exp ast.Expression, left interface{}, right interface{}) bool {
	accessExp, ok := exp.(*ast.AccessExpression)

	if !ok {
		t.Errorf("exp is not ast.AccessExpression. got=%T(%s)", exp, exp)
	}

	if !testLiteralExpression(t, accessExp.Left, left) {
		return false
	}

	if !testLiteralExpression(t, accessExp.Right, right) {
		return false
	}

	return true
}
