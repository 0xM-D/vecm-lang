package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestBooleanExpression(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("program has not got enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements)
	}

	testBooleanLiteral(t, stmt.Expression, true)

}
