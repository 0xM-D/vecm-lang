package parser

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
)

func TestForStatement(t *testing.T) {
	input := `for(int i = 0; i < 10; i+=1) { const int x = i * i; }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ForStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ForStatement. got=%T",
			program.Statements[0])
	}

	initialization, ok := stmt.Initialization.(*ast.TypedDeclarationStatement)
	if !ok {
		t.Fatalf("stmt.Intiialization is not ast.TypedDeclarationStatement. got=%T",
			stmt.Initialization)
	}

	testDeclarationStatement(t, &initialization.DeclarationStatement, "i", "0")

	condition, ok := stmt.Condition.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt.Condition is not ast.ExpressionStatement. got=%T", stmt.Condition)
	}

	testInfixExpression(t, condition.Expression, TestIdentifier{"i"}, "<", big.NewInt(10))

	afterThought, ok := stmt.AfterThought.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt.AfterThoguht is not ast.ExpressionStatement. got=%T", stmt.AfterThought)
	}

	testInfixExpression(t, afterThought.Expression, TestIdentifier{"i"}, "+=", big.NewInt(1))

	if len(stmt.Body.Statements) != 1 {
		t.Fatalf("for loop body is not 1 statements. got=%d\n",
			len(stmt.Body.Statements))
	}

	bodyStmt, ok := stmt.Body.Statements[0].(*ast.TypedDeclarationStatement)

	if !ok {
		t.Fatalf("stmt.Body.Statements[0] is not ast.TypedDeclarationStatement. got=%T", stmt.Body.Statements[0])
	}

	testDeclarationStatement(t, &bodyStmt.DeclarationStatement, "x", "(i * i)")

}

func TestEmptyForStatement(t *testing.T) {
	input := `for(;;) { }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ForStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ForStatement. got=%T",
			program.Statements[0])
	}

	if stmt.Initialization != nil {
		t.Fatalf("stmt.Intiialization is not nil. got=%T",
			stmt.Initialization)
	}

	if stmt.Condition != nil {
		t.Fatalf("stmt.Condition is not nil. got=%T",
			stmt.Initialization)
	}

	if stmt.AfterThought != nil {
		t.Fatalf("stmt.AfterThought is not nil. got=%T",
			stmt.Initialization)
	}

	if len(stmt.Body.Statements) != 0 {
		t.Fatalf("for loop body is not empty.\n")
	}

}

func TestForStatementNoBraces(t *testing.T) {
	input := `for(int i = 0; i < 10; i+=1) const int x = i * i;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ForStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ForStatement. got=%T",
			program.Statements[0])
	}

	initialization, ok := stmt.Initialization.(*ast.TypedDeclarationStatement)
	if !ok {
		t.Fatalf("stmt.Intiialization is not ast.TypedDeclarationStatement. got=%T",
			stmt.Initialization)
	}

	testDeclarationStatement(t, &initialization.DeclarationStatement, "i", "0")

	condition, ok := stmt.Condition.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt.Condition is not ast.ExpressionStatement. got=%T", stmt.Condition)
	}

	testInfixExpression(t, condition.Expression, TestIdentifier{"i"}, "<", big.NewInt(10))

	afterThought, ok := stmt.AfterThought.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt.AfterThoguht is not ast.ExpressionStatement. got=%T", stmt.AfterThought)
	}

	testInfixExpression(t, afterThought.Expression, TestIdentifier{"i"}, "+=", big.NewInt(1))

	if len(stmt.Body.Statements) != 1 {
		t.Fatalf("for loop body is not 1 statements. got=%d\n",
			len(stmt.Body.Statements))
	}

	bodyStmt, ok := stmt.Body.Statements[0].(*ast.TypedDeclarationStatement)

	if !ok {
		t.Fatalf("stmt.Body.Statements[0] is not ast.TypedDeclarationStatement. got=%T", stmt.Body.Statements[0])
	}

	testDeclarationStatement(t, &bodyStmt.DeclarationStatement, "x", "(i * i)")

}
