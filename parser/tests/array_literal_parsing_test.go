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
		{"new []int{1, 2 * 2, 3 + 3}", []string{"1", "(2 * 2)", "(3 + 3)"}},
		{"new []int{}", []string{}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
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
