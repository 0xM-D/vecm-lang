package evaluator_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestTypedDeclarationStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"int a = 5; a;", 5},
		{"string a = \"testmmm\"; a;", "testmmm"},
		{"const []int a = []int{1, 2, 3, 4}; let b = a; b;", []string{"1", "2", "3", "4"}},
		{"bool a = true; let b = !a; b;", false},
		{"const function(int, int)->int sum = fn(a: int, b: int) -> int { return a + b; }; sum", ExpectedFunction{
			"fn(a, b) {" + "\n" +
				"return (a + b);" + "\n" +
				"}",
			object.FunctionObjectType{
				ParameterTypes:  []object.ObjectType{object.IntegerKind, object.IntegerKind},
				ReturnValueType: object.IntegerKind,
			},
		}},
		{"function()->void sum = fn() -> void {}; sum", ExpectedFunction{
			"fn() {\n\n}",
			object.FunctionObjectType{
				ParameterTypes:  []object.ObjectType{},
				ReturnValueType: object.VoidKind,
			},
		}},
	}
	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, testEval(tt.input), int64(expected))
		case int64:
			testIntegerObject(t, testEval(tt.input), expected)
		case string:
			testStringObject(t, testEval(tt.input), expected)
		case bool:
			testBooleanObject(t, testEval(tt.input), expected)
		case []string:
			testArrayObject(t, testEval(tt.input), expected)
		case ExpectedFunction:
			testFunctionObject(t, testEval(tt.input), expected)
		}
	}
}
