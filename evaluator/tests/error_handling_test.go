package evaluator_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"operator + not defined on types int and bool",
		},
		{
			"5 + true; 5;",
			"operator + not defined on types int and bool",
		},
		{
			"-true",
			"unknown operator: -bool",
		},
		{
			"true + false;",
			"operator + not defined on types bool and bool",
		},
		{
			"5; true + false; 5",
			"operator + not defined on types bool and bool",
		},
		{
			"if (10 > 1) { true + false; }",
			"operator + not defined on types bool and bool",
		},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}
					return 1;
				}
			`,
			"operator + not defined on types bool and bool",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"operator - not defined on types string and string",
		},
		{
			`{"name": "Monkey"}[fn(x:string)->string { x }];`,
			"unusable as hash key: function(string) -> string",
		},
		{
			`int a = "fasdf"`,
			"Expression of type string cannot be assigned to int",
		},
		{
			`a := "fasdf"; bool c = a;`,
			"Expression of type string cannot be assigned to bool",
		},
		{
			`a := "fasdf"; string a = a;`,
			"Identifier with name a already exists.",
		},
		{
			`const abcc = "fasdf"; abcc = "fasdfsd";`,
			"Cannot assign to const variable",
		},
		{
			`const int a = 3; a = 5*6;`,
			"Cannot assign to const variable",
		},
		{
			`foo += 3`,
			"identifier not found: foo",
		},
		{
			`a := 3; a += "test"`,
			"operator += not defined on types int and string",
		},
		{
			`bool a = true; a += true`,
			"operator += not defined on types bool and bool",
		},
		{
			`const int a = 3; a += 1`,
			"Cannot assign to const variable",
		},
		{
			`fun := fn()->void {}; fun(1)`,
			"Incorrect parameter count for function() -> void fun. expected=0, got=1",
		},
		{
			`fun := fn(a: int, b: int)->int { return a + b; }; fun()`,
			"Incorrect parameter count for function(int, int) -> int fun. expected=2, got=0",
		},
		{
			`[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].deleteee(1, 3)`,
			"Member deleteee does not exist on int[]",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)

		if evaluated == nil || !object.IsError(evaluated) {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		errObj := evaluated.(*object.Error)

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}
