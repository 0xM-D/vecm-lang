package evaluator

import (
	"math/big"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected *big.Int
	}{
		{"return 10;", big.NewInt(10)},
		{"return 10; 9;", big.NewInt(10)},
		{"return 2 * 5; 9;", big.NewInt(10)},
		{"9; return 2 * 5; 9;", big.NewInt(10)},
		{`	if (10 > 1) {
				if (10 > 1) {
					return 10;
				}
				return 1;
			}`, big.NewInt(10)},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
