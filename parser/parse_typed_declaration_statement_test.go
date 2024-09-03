package parser

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
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
		{"function(int) -> int c = fn(b:int)->int { return b * 2 };", "c", "function(int)->int", "fn(b:int)->int{return (b * 2);}"},
		{"function(int, int)->int sum = fn(a: int, b: int) -> int { return a + b; }", "sum", "function(int, int)->int", "fn(a:int,b:int)->int{return (a + b);}"},
		{"map{ int -> int } d = new map{ int -> int }{1: 2, 2: 3};", "d", "map{ int -> int }", "new map{ int -> int }{1: 2, 2: 3}"},
		{"[]int e = new []int{1, 2, 3, 4, 5};", "e", "[]int", "new []int{1, 2, 3, 4, 5}"},
		{"[][]int e = new [][]int{new []int{1, 2, 3, 4, 5}};", "e", "[][]int", "new [][]int{new []int{1, 2, 3, 4, 5}}"},
		{`[]map{ string -> int } d = new []map{ string -> int }{new map{ int -> int }{"foo": 2, "bar": 3}};`, "d", "[]map{ string -> int }", `new []map{ string -> int }{new map{ int -> int }{"foo": 2, "bar": 3}}`},
		{`[]map{ string -> []int } d = new []map{ string -> []int }{new map{ string -> []int }{"foo": new []int{1, 2}, "bar": new []int{3, 4}}};`, "d", "[]map{ string -> []int }", `new []map{ string -> []int }{new map{ string -> []int }{"foo": new []int{1, 2}, "bar": new []int{3, 4}}}`},
		{`string f = "string value";`, "f", "string", `"string value"`},
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

		decl, ok := program.Statements[0].(*ast.TypedDeclarationStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.TypedDeclarationStatement. got=%T", program.Statements[0])
		}

		if decl.Name.Value != tt.expectedName {
			t.Fatalf("decl.Name is not %q. got=%q", tt.expectedName, decl.Name.Value)
		}

		if decl.Type == nil {
			t.Fatalf("decl.Type is null")
		}

		if decl.Type.String() != tt.expectedType {
			t.Fatalf("decl.Type is not %s. got=%s", tt.expectedType, decl.Type.String())
		}

		if decl.Value.String() != tt.expectedValue {
			t.Fatalf("decl.Value is not %s. got=%s", tt.expectedValue, decl.Value.String())
		}
	}
}

func TestConstDeclarationStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue string
	}{
		{"const a = 10;", "a", "10"},
		{"const bool a = true;", "a", "true"},
		{"const c = fn(b: int)->int { return b * 2 };", "c", "fn(b:int)->int{return (b * 2);}"},
		{"const a = new map{int -> int}{1: 2, 2: 3};", "a", "new map{ int -> int }{1: 2, 2: 3}"},
		{"const array e = new []int{1, 2, 3, 4, 5};", "e", "new []int{1, 2, 3, 4, 5}"},
		{`const string f = "string value";`, "f", `"string value"`},
		{`const f = "string value";`, "f", `"string value"`},
		{`const a = 10;`, "a", "10"},
		{`const arr = new []int{1, 2, 3, 4, 5};`, "arr", "new []int{1, 2, 3, 4, 5}"},
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
