package evaluator

import (
	"testing"
)

func TestRuntimeErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"operator + not defined on types int64 and bool",
		},
		{
			"5 + true; 5;",
			"operator + not defined on types int64 and bool",
		},
		{
			"-true",
			"operator - not defined on type bool",
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
			`int a = "fasdf"`,
			"expression of type string cannot be assigned to int64",
		},
		{
			`a := "fasdf"; bool c = a;`,
			"expression of type string cannot be assigned to bool",
		},
		{
			`a := "fasdf"; string a = a;`,
			"identifier with name a already exists",
		},
		{
			`const abcc = "fasdf"; abcc = "fasdfsd";`,
			"cannot assign to const variable",
		},
		{
			`const int32 a = 3; a = 5*6;`,
			"cannot assign to const variable",
		},
		{
			`foo += 3`,
			"identifier not found: foo",
		},
		{
			`a := 3; a += "test"`,
			"operator += not defined on types int64 and string",
		},
		{
			`bool a = true; a += true`,
			"operator += not defined on types bool and bool",
		},
		{
			`const int64 a = 3; a += 1`,
			"cannot assign to const variable",
		},
		{
			`fun := fn()->void {}; fun(1)`,
			"incorrect parameter count for function() -> void fun. expected=0, got=1",
		},
		{
			`fun := fn(a: int64, b: int64)->int64 { return a + b; }; fun()`,
			"incorrect parameter count for function(int64, int64) -> int64 fun. expected=2, got=0",
		},
		{
			`new []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.deleteee(1, 3)`,
			"member deleteee does not exist on []int64",
		},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)

		if err == nil {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if err.Error() != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, err.Error())
		}
	}
}
