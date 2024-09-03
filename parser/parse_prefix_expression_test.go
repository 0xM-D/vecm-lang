package parser

import (
	"math/big"
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
)

func TestPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5", "!", big.NewInt(5)},
		{"-15", "-", big.NewInt(15)},
		{"!true;", "!", true},
		{"!false;", "!", false},
		{"~1;", "~", big.NewInt(1)},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpresssionStatement. got=%T", program.Statements[0])
		}

		expr, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if expr.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, expr.Operator)
		}

		if !testLiteralExpression(t, expr.Right, tt.value) {
			return
		}
	}
}
