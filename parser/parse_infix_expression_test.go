package parser_test

import (
	"math/big"
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", big.NewInt(5), "+", big.NewInt(5)},
		{"5 - 5;", big.NewInt(5), "-", big.NewInt(5)},
		{"5 * 5;", big.NewInt(5), "*", big.NewInt(5)},
		{"5 / 5;", big.NewInt(5), "/", big.NewInt(5)},
		{"5 > 5;", big.NewInt(5), ">", big.NewInt(5)},
		{"5 < 5;", big.NewInt(5), "<", big.NewInt(5)},
		{"5 == 5;", big.NewInt(5), "==", big.NewInt(5)},
		{"5 != 5;", big.NewInt(5), "!=", big.NewInt(5)},
		{"foobar + barfoo;", TestIdentifier{"foobar"}, "+", TestIdentifier{"barfoo"}},
		{"foobar - barfoo;", TestIdentifier{"foobar"}, "-", TestIdentifier{"barfoo"}},
		{"foobar * barfoo;", TestIdentifier{"foobar"}, "*", TestIdentifier{"barfoo"}},
		{"foobar / barfoo;", TestIdentifier{"foobar"}, "/", TestIdentifier{"barfoo"}},
		{"foobar > barfoo;", TestIdentifier{"foobar"}, ">", TestIdentifier{"barfoo"}},
		{"foobar < barfoo;", TestIdentifier{"foobar"}, "<", TestIdentifier{"barfoo"}},
		{"foobar >= barfoo;", TestIdentifier{"foobar"}, ">=", TestIdentifier{"barfoo"}},
		{"foobar <= barfoo;", TestIdentifier{"foobar"}, "<=", TestIdentifier{"barfoo"}},
		{"foobar == barfoo;", TestIdentifier{"foobar"}, "==", TestIdentifier{"barfoo"}},
		{"foobar != barfoo;", TestIdentifier{"foobar"}, "!=", TestIdentifier{"barfoo"}},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
		{"true && false", true, "&&", false},
		{"false || false", false, "||", false},
		{"foobar & barfoo;", TestIdentifier{"foobar"}, "&", TestIdentifier{"barfoo"}},
		{"foobar | barfoo;", TestIdentifier{"foobar"}, "|", TestIdentifier{"barfoo"}},
		{"foobar ^ barfoo;", TestIdentifier{"foobar"}, "^", TestIdentifier{"barfoo"}},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not an ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		if !testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * new []int{1, 2, 3, 4}[b * c] * d",
			"((a * (new []int{1, 2, 3, 4}[(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * new []int{1, 2}[1])",
			"add((a * (b[2])), (b[1]), (2 * (new []int{1, 2}[1])))",
		},
		{
			"(a < b && a > c)",
			"((a < b) && (a > c))",
		},
		{
			"(a < b && a > c || c != d)",
			"(((a < b) && (a > c)) || (c != d))",
		},
		{
			"a * b < c + d * e && a == c || -d != e / f",
			"((((a * b) < (c + (d * e))) && (a == c)) || ((-d) != (e / f)))",
		},
		{
			"~a & b | c == 13 * 4",
			"(((~a) & b) | (c == (13 * 4)))",
		},
		{
			"(~a & b | c) == 13 * 4",
			"((((~a) & b) | c) == (13 * 4))",
		},
		{
			"a << 3 == 0",
			"((a << 3) == 0)",
		},
		{
			"a >> b && 3",
			"((a >> b) && 3)",
		},
		{
			"true && false ? true & false : 1 + 2",
			"((true && false) ? (true & false) : (1 + 2))",
		},
		{
			"false ? (false ? 1 : 2) : true ? 3 : 4",
			"(false ? (false ? 1 : 2) : (true ? 3 : 4))",
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
