package parser

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
)

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("stmt is not ast.IfStatement. got=%T", stmt)
	}

	if !testInfixExpression(t, stmt.Condition, TestIdentifier{"x"}, "<", TestIdentifier{"y"}) {
		return
	}

	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("stmt.Consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}

	consequence, ok := stmt.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Consequence.Statements[0])
	}
	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}
	if stmt.Alternative != nil {
		t.Errorf("stmt.Alternative.Statements was not nil. got=%+v", stmt.Alternative)
	}

}

func TestIfExpressionNoBrace(t *testing.T) {
	input := `if x < y x else y`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("stmt is not ast.IfStatement. got=%T", stmt)
	}

	if !testInfixExpression(t, stmt.Condition, TestIdentifier{"x"}, "<", TestIdentifier{"y"}) {
		return
	}

	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("stmt.Consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}

	consequence, ok := stmt.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Consequence.Statements[0])
	}
	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if len(stmt.Alternative.Statements) != 1 {
		t.Errorf("stmt.Alternative is not 1 statements. got=%d\n",
			len(stmt.Alternative.Statements))
	}

	alternative, ok := stmt.Alternative.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", stmt.Consequence.Statements[0])
	}
	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}

}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("stmt is not ast.IfStatement. got=%T", stmt)
	}

	if !testInfixExpression(t, stmt.Condition, TestIdentifier{"x"}, "<", TestIdentifier{"y"}) {
		return
	}

	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("stmt.Consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}

	consequence, ok := stmt.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Consequence.Statements[0])
	}
	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if stmt.Alternative == nil {
		t.Fatalf("exp.Alternative.Statements was nil")
	}

	if len(stmt.Alternative.Statements) != 1 {
		t.Errorf("stmt.Alternative is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}

	alternative, ok := stmt.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Alternative.Statements[0])
	}
	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}

}
