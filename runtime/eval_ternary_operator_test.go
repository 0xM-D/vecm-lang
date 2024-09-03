package runtime_test

import (
	"math/big"
	"testing"
)

func TestTernaryOperator(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{" true ? 1 : 2", big.NewInt(1)},
		{"foo:= 64; bar := 88; false ? foo : bar", big.NewInt(88)},
		{` 1 < 2 && 3 <= 3 ? "brr" : "gzz" `, "brr"},
		{` true ? 1 + 2 : 5`, big.NewInt(3)},
		{` false ? (false ? 1 : 2) : true ? 3 : 4`, big.NewInt(3)},
	}

	for _, tt := range tests {
		result, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testLiteralObject(t, result, tt.expectedValue)
	}
}
