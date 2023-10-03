package parser_tests

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestAccessExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		rightValue interface{}
	}{
		{"5.toString", big.NewInt(5), TestIdentifier{"toString"}},
		{"array.size", TestIdentifier{"array"}, TestIdentifier{"size"}},
		{"[]int{1, 2, 3, 4}.size", []interface{}{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4)}, TestIdentifier{"size"}},
		{"str.length", TestIdentifier{"str"}, TestIdentifier{"length"}},
		{`"ABCDEFG".length`, "ABCDEFG", TestIdentifier{"length"}},
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

		if !testAccessExpression(t, stmt.Expression, tt.leftValue, tt.rightValue) {
			return
		}
	}
}
