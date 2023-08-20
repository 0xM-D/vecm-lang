package evaluator_tests

import (
	"testing"
)

func TestEvalFloatExpression(t *testing.T) {

	tests := []struct {
		input    string
		expected interface{}
	}{
		{"5.0", 5.0},
		{"1.1f", float32(1.1)},
		{"-5f", float32(-5.0)},
		{"-10.22233344f", -float32(10.22233344)},
		{"5.0 + 5.0f + .5 + 5 - 10", 5.5},
		{"2.0 * 2f * 2.0f * 2 * 2", 32.0},
		{"-50 + 100 + -50", 0},
		{"5f * 2 + 10f", float32(20)},
		{"5 + 2f * 10", float32(25)},
		{"20f + 2.0 * -10f", 0.0},
		{"51 / 2 * 2f + 10f", float32(60)},
		{"2 * (5f + 10.0)", 30.0},
		{"3f * 3f * 3f + 10", float32(37.0)},
		{"3 * (3f * 3) + 10", float32(37.0)},
		{"(5.0f + 10 * 2 + 11f / 3) * 2 + -10", float32(47.333333333333336)},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumber(t, evaluated, tt.expected)
	}
}
