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
		{"function(int) -> int c = fn(b:int)->int { return b * 2 };", "c", "function(int)->int", "fn(b:int)->int{return (b * 2);}"},
		{"function(int, int)->int sum = fn(a: int, b: int) -> int { return a + b; }", "sum", "function(int, int)->int", "fn(a:int,b:int)->int{return (a + b);}"},
		{"map{ int -> int } d = {1: 2, 2: 3};", "d", "map{ int -> int }", "{1:2, 2:3}"},
		{"int[] e = [1, 2, 3, 4, 5];", "e", "int[]", "[1, 2, 3, 4, 5]"},
		{"int[][] e = [[1, 2, 3, 4, 5]];", "e", "int[][]", "[[1, 2, 3, 4, 5]]"},
		{`map{ string -> int }[] d = [{"foo": 2, "bar": 3}];`, "d", "map{ string -> int }[]", `[{foo:2, bar:3}]`},
		{`map{ string -> int[] }[] d = [{"foo": [1, 2], "bar": [3, 4]}];`, "d", "map{ string -> int[] }[]", `[{foo:[1, 2], bar:[3, 4]}]`},
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
