package parser_test

import (
	"math/big"
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestNewArrayArrayExpression(t *testing.T) {
	tests := []struct {
		input            string
		expectedElements []string
	}{
		{"new []int{1, 2 * 2, 3 + 3}", []string{"1", "(2 * 2)", "(3 + 3)"}},
		{"new []int{}", []string{}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, isExpressionStatement := program.Statements[0].(*ast.ExpressionStatement)
		if !isExpressionStatement {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.NewExpression)
		if !ok {
			t.Fatalf("exp not ast.NewExpression. got=%T", stmt.Expression)
		}

		_, ok = expr.Type.(ast.ArrayType)

		if !ok {
			t.Fatalf("expr.Type not ast.ArrayType. got=%T", expr.Type)
		}

		if len(expr.InitializationList) != len(tt.expectedElements) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(expr.InitializationList), len(tt.expectedElements))
		}

		for index, expected := range tt.expectedElements {
			got := expr.InitializationList[index].String()
			if got != expected {
				t.Errorf("Array element doesn't match expected. expected=%s got=%s", expected, got)
			}
		}
	}
}

func TestNewHashExpression(t *testing.T) {
	input := `new map{ string -> int }{"one": 1, "two": 2, "three": 3}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, isExpressionStatement := program.Statements[0].(*ast.ExpressionStatement)
	if !isExpressionStatement {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := stmt.Expression.(*ast.NewExpression)
	if !ok {
		t.Fatalf("exp is not ast.NewExpression. got=%T", stmt.Expression)
	}

	_, ok = hash.Type.(ast.HashType)
	if !ok {
		t.Errorf("newExpression.Type is not ast.HashType. got=%T", hash.Type)
	}

	if len(hash.InitializationList) != 3 {
		t.Errorf("hash.InitializationList has wrong length. got=%d", len(hash.InitializationList))
	}

	expected := map[string]*big.Int{
		"one":   big.NewInt(1),
		"two":   big.NewInt(2),
		"three": big.NewInt(3),
	}

	for _, expr := range hash.InitializationList {
		pair, isPairExpression := expr.(*ast.PairExpression)
		if !isPairExpression {
			t.Errorf("expr is not ast.PairExpression. got=%T", expr)
		}

		key, isStringLiteral := pair.Left.(*ast.StringLiteral)
		if !isStringLiteral {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
		}

		expectedValue := expected[key.Value]
		testIntegerLiteral(t, pair.Right, expectedValue)
	}
}

func TestNewHashLiteralExpressionEmpty(t *testing.T) {
	input := "new map{ int -> []int }{}"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, isExpressionStatement := program.Statements[0].(*ast.ExpressionStatement)
	if !isExpressionStatement {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := stmt.Expression.(*ast.NewExpression)

	if !ok {
		t.Fatalf("exp is not ast.NewExpression. got=%T", stmt.Expression)
	}

	if len(hash.InitializationList) != 0 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.InitializationList))
	}
}

func TestNewHashLiteralExpressionWithExpressions(t *testing.T) {
	input := `new map{ string -> int }{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, isExpressionStatement := program.Statements[0].(*ast.ExpressionStatement)
	if !isExpressionStatement {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, isNewExpression := stmt.Expression.(*ast.NewExpression)

	if !isNewExpression {
		t.Fatalf("exp is not ast.NewExpression. got=%T", stmt.Expression)
	}

	if len(hash.InitializationList) != 3 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.InitializationList))
	}

	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, big.NewInt(0), "+", big.NewInt(1))
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, big.NewInt(10), "-", big.NewInt(8))
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, big.NewInt(15), "/", big.NewInt(5))
		},
	}

	for _, expr := range hash.InitializationList {
		pair, isPairExpression := expr.(*ast.PairExpression)
		if !isPairExpression {
			t.Errorf("expr is not ast.PairExpression. got=%T", expr)
		}

		literal, isStringLiteral := pair.Left.(*ast.StringLiteral)
		if !isStringLiteral {
			t.Errorf("key is not ast.StringLiteral. got=%T", pair.Left)
			continue
		}

		testFunc, ok := tests[literal.Value]
		if !ok {
			t.Errorf("No test function for key %q found", literal.String())
			continue
		}

		testFunc(pair.Right)
	}
}
