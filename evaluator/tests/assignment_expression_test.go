package evaluator_tests

import (
	"math/big"
	"testing"
)

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
