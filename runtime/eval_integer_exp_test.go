package runtime_test

import (
	"math/big"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected *big.Int
	}{
		{"5", big.NewInt(5)},
		{"10", big.NewInt(10)},
		{"-5", big.NewInt(-5)},
		{"-10", big.NewInt(-10)},
		{"5 + 5 + 5 + 5 - 10", big.NewInt(10)},
		{"2 * 2 * 2 * 2 * 2", big.NewInt(32)},
		{"-50 + 100 + -50", big.NewInt(0)},
		{"5 * 2 + 10", big.NewInt(20)},
		{"5 + 2 * 10", big.NewInt(25)},
		{"20 + 2 * -10", big.NewInt(0)},
		{"50 / 2 * 2 + 10", big.NewInt(60)},
		{"2 * (5 + 10)", big.NewInt(30)},
		{"3 * 3 * 3 + 10", big.NewInt(37)},
		{"3 * (3 * 3) + 10", big.NewInt(37)},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", big.NewInt(50)},
		{"1 << 1", big.NewInt(2)},
		{"1 << 62", big.NewInt(4611686018427387904)},
		{"1 << 63 >> 1", big.NewInt(1 << 63 >> 1)},
		{"1 << 64 >> 2", big.NewInt(0)},
		{"1 >> 1", big.NewInt(0)},
		{"256 >> 2", big.NewInt(64)},
		{"3 >> 1", big.NewInt(1)},
		{"3 << 1", big.NewInt(6)},
		{"1 | 3", big.NewInt(3)},
		{"4097 | 272", big.NewInt(4369)},
		{"0 | 272", big.NewInt(272)},
		{"0 | 0", big.NewInt(0)},
		{"0 & 0", big.NewInt(0)},
		{"0 & 1", big.NewInt(0)},
		{"4097 & 272", big.NewInt(0)},
		{"7 & 3", big.NewInt(3)},
		{"~0", big.NewInt(^0)},
		{"((1 << 10) - 1) ^ (1 << 8)", big.NewInt(int64((1<<10)-1) ^ (1 << 8))},
		{"~123", big.NewInt(^123)},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testIntegerObject(t, evaluated, tt.expected)
	}
}
