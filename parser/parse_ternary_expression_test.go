package parser_test

import (
	"math/big"
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestTernaryOperator(t *testing.T) {
	tests := []struct {
		input        string
		condition    interface{}
		valueIfTrue  interface{}
		valueIfFalse interface{}
	}{
		{" true ? 1 : 2", true, big.NewInt(1), big.NewInt(2)},
		{" false ? foo : bar", false, TestIdentifier{"foo"}, TestIdentifier{"bar"}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Program does not have one statement. got=%d", len(program.Statements))
		}

		expressionStatement, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not an expressionStatement. got=%T", program.Statements[0])
		}

		ternaryExpression, ok := expressionStatement.Expression.(*ast.TernaryExpression)

		if !ok {
			t.Fatalf("expressionStatement.Expression is not ternaryExpression. got=%T", expressionStatement.Expression)
		}

		if !testLiteralExpression(t, ternaryExpression.Condition, tt.condition) {
			t.Fatalf("Ternary expression condition is different than expected. want=%s got=%s",
				ternaryExpression.Condition.String(),
				tt.condition,
			)
		}

		if !testLiteralExpression(t, ternaryExpression.ValueIfTrue, tt.valueIfTrue) {
			t.Fatalf("Ternary expression valueIfTrue is different than expected. want=%s got=%s",
				ternaryExpression.ValueIfTrue.String(),
				tt.valueIfTrue,
			)
		}

		if !testLiteralExpression(t, ternaryExpression.ValueIfFalse, tt.valueIfFalse) {
			t.Fatalf("Ternary expression falueIfFalse is different than expected. want=%s got=%s",
				ternaryExpression.ValueIfFalse.String(),
				tt.valueIfFalse,
			)
		}
	}
}
