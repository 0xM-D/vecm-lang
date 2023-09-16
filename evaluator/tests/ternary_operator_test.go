package evaluator_tests

import (
	"testing"
)

func TestTernaryOperator(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{" true ? 1 : 2", 1},
		{"foo:= 64; bar := 88; false ? foo : bar", 88},
		{` 1 < 2 && 3 <= 3 ? "brr" : "gzz" `, "brr"},
	}

	for _, tt := range tests {

		testLiteralObject(t, testEval(tt.input), tt.expectedValue)

	}
}
