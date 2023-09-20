package evaluator_tests

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestTypedDeclarationStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"int64 a = 5; a;", big.NewInt(5)},
		{"string a = \"testmmm\"; a;", "testmmm"},
		{"const int8[] a = [1, 2, 3, 4]; let b = a; b;", []string{"1", "2", "3", "4"}},
		{"bool a = true; let b = !a; b;", false},
		{"const function(int64, int64)->int64 sum = fn(a: int64, b: int64) -> int64 { return a + b; }; sum", ExpectedFunction{
			"fn(a, b) {" + "\n" +
				"return (a + b);" + "\n" +
				"}",
			object.FunctionObjectType{
				ParameterTypes:  []object.ObjectType{object.Int64Kind, object.Int64Kind},
				ReturnValueType: object.Int64Kind,
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
		case *big.Int:
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
