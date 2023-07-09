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
			"type mismatch: int + bool",
		},
		{
			"5 + true; 5;",
			"type mismatch: int + bool",
		},
		{
			"-true",
			"unknown operator: -bool",
		},
		{
			"true + false;",
			"unknown operator: bool + bool",
		},
		{
			"5; true + false; 5",
			"unknown operator: bool + bool",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: bool + bool",
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
			"unknown operator: bool + bool",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: string - string",
		},
		{
			`{"name": "Monkey"}[fn(x) { x }];`,
			"unusable as hash key: function",
		},
		{
			`int a = "fasdf"`,
			"Expression of type string cannot be assigned to int",
		},
		{
			`a := "fasdf"; bool c = a;`,
			"Expression of type string cannot be assigned to bool",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}
