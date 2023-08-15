package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestParsingArrayLiterals(t *testing.T) {
	tests := []struct {
		input            string
		expectedElements []string
	}{
		{"[1, 2 * 2, 3 + 3]", []string{"1", "(2 * 2)", "(3 + 3)"}},
		{"[]", []string{}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		array, ok := stmt.Expression.(*ast.ArrayLiteral)

		if !ok {
			t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
		}

		if len(array.Elements) != len(tt.expectedElements) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(array.Elements), len(tt.expectedElements))
		}

		for index, expected := range tt.expectedElements {
			got := array.Elements[index].String()
			if got != expected {
				t.Errorf("Array element doesn't match expected. expected=%s got=%s", expected, got)
			}
		}
	}

}
