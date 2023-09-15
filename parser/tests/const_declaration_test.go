package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestConstDeclaration(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue string
	}{
		{"const a = 10;", "a", "10"},
		{"const bool a = true;", "a", "true"},
		{"const c = fn(b: int)->int { return b * 2 };", "c", "fn(b:int)->int{return (b * 2);}"},
		{"const a = {1: 2, 2: 3};", "a", "{1:2, 2:3}"},
		{"const array e = [1, 2, 3, 4, 5];", "e", "[1, 2, 3, 4, 5]"},
		{`const string f = "string value";`, "f", "string value"},
		{`const f = "string value";`, "f", "string value"},
		{`const a = 10;`, "a", "10"},
		{`const arr = [1, 2, 3, 4, 5];`, "arr", "[1, 2, 3, 4, 5]"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		switch decl := program.Statements[0].(type) {
		case *ast.TypedDeclarationStatement:
			testDeclarationStatement(t, &decl.DeclarationStatement, tt.expectedName, tt.expectedValue)
		case *ast.AssignmentDeclarationStatement:
			testDeclarationStatement(t, &decl.DeclarationStatement, tt.expectedName, tt.expectedValue)
		default:
			t.Fatalf("program.Statements[0] is not ast.DeclarationStatement. got=%T", program.Statements[0])
		}

	}
}

func testDeclarationStatement(t *testing.T, decl *ast.DeclarationStatement, expectedName string, expectedValue string) {
	if decl.Name.Value != expectedName {
		t.Fatalf("decl.Name is not %q. got=%q", expectedName, decl.Name.Value)
	}

	if decl.Value.String() != expectedValue {
		t.Fatalf("decl.Value is not %s. got=%s", expectedValue, decl.Value.String())
	}

}
