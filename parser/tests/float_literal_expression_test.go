package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestFloat32Literal(t *testing.T) {
	input := "5f; 5.5f; .5f; 0.55f;"
	expected := []float32{5, 5.5, .5, 0.55}

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("program does not have 4 statements. got=%d", len(program.Statements))
	}

	for i, statement := range program.Statements {
		stmt, ok := statement.(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[%d] is not an *ast.ExpressionStatement. got=%T", i, program.Statements[0])
		}

		testFloat32Literal(t, stmt.Expression, expected[i])
	}

}

func TestFloat64Literal(t *testing.T) {
	input := "5.0; 5.5; .5; 0.55;"
	expected := []float64{5, 5.5, .5, 0.55}

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("program does not have 4 statements. got=%d", len(program.Statements))
	}

	for i, statement := range program.Statements {
		stmt, ok := statement.(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[%d] is not an *ast.ExpressionStatement. got=%T", i, program.Statements[0])
		}

		testFloat64Literal(t, stmt.Expression, expected[i])
	}

}
