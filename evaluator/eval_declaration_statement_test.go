package evaluator

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestAssignmentDeclaration(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"a := 5; a;", 5},
		{"a := 5 * 5; a;", 25},
		{"a := 5; let b = a; b;", 5},
		{"a := 5; let b = a; let c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), big.NewInt(tt.expected))
	}
}

func TestAsignmentExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let a = 5; a = 3;", big.NewInt(3)},
		{"a := 5 * 5; a += 20;", big.NewInt(45)},
		{"a := 3; b := a; a += b", big.NewInt(6)},
		{"a := 5; a -= 1", big.NewInt(4)},
		{"a := 5; a *= 2", big.NewInt(10)},
		{"a := 50; a /= 5", big.NewInt(10)},
		{`a := "a"; a += "bc"`, "abc"},
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
		}
	}
}

func TestTypedDeclarationStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"int64 a = 5; a;", big.NewInt(5)},
		{"string a = \"testmmm\"; a;", "testmmm"},
		{"const []int a = new []int{1, 2, 3, 4}; let b = a; b;", []string{"1", "2", "3", "4"}},
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

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected *big.Int
	}{
		{"let a = 5; a;", big.NewInt(5)},
		{"let a = 5 * 5; a;", big.NewInt(25)},
		{"let a = 5; let b = a; b;", big.NewInt(5)},
		{"let a = 5; let b = a; let c = a + b + 5; c;", big.NewInt(15)},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
