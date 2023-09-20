package evaluator_tests

import (
	"math/big"
	"testing"
)

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
