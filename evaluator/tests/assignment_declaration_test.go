package evaluator_tests

import (
	"math/big"
	"testing"
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
