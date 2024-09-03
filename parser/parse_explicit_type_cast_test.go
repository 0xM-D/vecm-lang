package parser

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
)

func TestExplicitTypeCast(t *testing.T) {
	tests := []struct {
		input        string
		expectedLeft string
		expectedType string
	}{
		{"8 as uint64", "8", "uint64"},
		{"(6.0f * 2) as float64", "(6.0f * 2)", "float64"},
		{"new []int16{1, 2, 3} as []int64", "new []int16{1, 2, 3}", "[]int64"},
		{`new map{ string -> string }{ 1: "2", 3: "4"} as map{ string -> string }`, `new map{ string -> string }{1: "2", 3: "4"}`, "map{ string -> string }"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		exprStmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		typeCastExpr, ok := exprStmt.Expression.(*ast.TypeCastExpression)

		if !ok {
			t.Fatalf("exprStmt is not ast.TypeCastExpression. got=%T", exprStmt.Expression)
		}

		if typeCastExpr.Left.String() != tt.expectedLeft {
			t.Fatalf("typeCastExpr.Left is not %s. got=%s", tt.expectedLeft, typeCastExpr.Left.String())
		}

		if typeCastExpr.Type.String() != tt.expectedType {
			t.Fatalf("typeCastExpr.Type is not %s. got=%s", tt.expectedType, typeCastExpr.Type.String())
		}
	}
}
