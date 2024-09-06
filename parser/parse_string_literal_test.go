package parser_test

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestStringLiteral(t *testing.T) {
	input := `"test string"`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not have 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	str, ok := stmt.Expression.(*ast.StringLiteral)

	if !ok {
		t.Fatalf("stmt.Expression not *ast.StringLiteral. got=%T", stmt.Expression)
	}

	if str.Value != "test string" {
		t.Fatalf("str.Value not %q. got=%q", "test string", str.Value)
	}
}
