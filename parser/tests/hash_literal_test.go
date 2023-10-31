package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestParsingHashLiteralsStringKeys(t *testing.T) {
	input := `new map{ string -> int }{"one": 1, "two": 2, "three": 3}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
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

	expected := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for _, expr := range hash.InitializationList {
		pair, ok := expr.(*ast.PairExpression)

		if !ok {
			t.Errorf("expr is not ast.PairExpression. got=%T", expr)
		}

		key, ok := pair.Left.(*ast.StringLiteral)

		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
		}

		expectedValue := expected[key.Value]
		testIntegerLiteral(t, pair.Right, expectedValue)
	}
}

func TestParsingEmptyHashLiteral(t *testing.T) {
	input := "new map{ int -> []int }{}"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.NewExpression)

	if !ok {
		t.Fatalf("exp is not ast.NewExpression. got=%T", stmt.Expression)
	}

	if len(hash.InitializationList) != 0 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.InitializationList))
	}

}

func TestParsingHashLiteralsWithExpressions(t *testing.T) {
	input := `new map{ string -> int }{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.NewExpression)

	if !ok {
		t.Fatalf("exp is not ast.NewExpression. got=%T", stmt.Expression)
	}

	if len(hash.InitializationList) != 3 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.InitializationList))
	}

	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, 0, "+", 1)
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, 10, "-", 8)
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, 15, "/", 5)
		},
	}

	for _, expr := range hash.InitializationList {
		pair, ok := expr.(*ast.PairExpression)

		if !ok {
			t.Errorf("expr is not ast.PairExpression. got=%T", expr)
		}

		literal, ok := pair.Left.(*ast.StringLiteral)

		if !ok {
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
