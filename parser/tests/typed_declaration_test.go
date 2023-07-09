package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestTypedDeclaration(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedType  string
		expectedValue string
	}{
		{"int a = 10;", "a", "int", "10"},
		{"bool b = true;", "b", "bool", "true"},
		{"function c = fn(b) { return b * 2 };", "c", "function", "fn(b)return (b * 2);"},
		{"map d = {1: 2, 2: 3};", "d", "map", "{1:2, 2:3}"},
		{"array e = [1, 2, 3, 4, 5];", "e", "array", "[1, 2, 3, 4, 5]"},
		{`string f = "string value";`, "f", "string", "string value"},
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

		decl, ok := program.Statements[0].(*ast.TypedDeclarationStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.TypedDeclarationStatement. got=%T", program.Statements[0])
		}

		if decl.Name.Value != tt.expectedName {
			t.Fatalf("decl.Name is not %q. got=%q", tt.expectedName, decl.Name.Value)
		}

		if decl.Type.String() != tt.expectedType {
			t.Fatalf("decl.Type is not %s. got=%s", tt.expectedType, decl.Type.String())
		}

		if decl.Value.String() != tt.expectedValue {
			t.Fatalf("decl.Value is not %s. got=%s", tt.expectedValue, decl.Value.String())
		}

	}
}
