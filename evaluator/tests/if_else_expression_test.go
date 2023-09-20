package evaluator_tests

import (
	"math/big"
	"testing"
)

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", big.NewInt(10)},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", big.NewInt(10)},
		{"if (1 < 2) { 10 }", big.NewInt(10)},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", big.NewInt(20)},
		{"if (1 < 2) { 10 } else { 20 }", big.NewInt(10)},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(*big.Int)
		if ok {
			testIntegerObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}

}
