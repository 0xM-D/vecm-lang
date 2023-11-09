package evaluator

import (
	"math/big"
	"testing"
)

func TestTypeBuiltins(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let arr = new []int{1, 2, 3}; arr.size();", big.NewInt(3)},
		{"let arr = new []int{}; arr.size();", big.NewInt(0)},
		{"new []int{1, 2, 3, 4, 5}.size();", big.NewInt(5)},
		{`str := "abcdef"; str.length();`, big.NewInt(6)},
		{`const string str = ""; str.length();`, big.NewInt(0)},
		{`"bleh".length()`, big.NewInt(4)},
		{"1.toString()", "1"},
		{"(123*456).toString()", "56088"},
		{"let arr = new []int{1, 2, 3}; arr.push(4).size()", big.NewInt(4)},
		{"new []int{}.push(1).size();", big.NewInt(1)},
		{"new []int{1, 2, 3}.delete(1, 5).size();", big.NewInt(1)},
		{`new []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}.delete(1, 3);`, []string{"1", "5", "6", "7", "8", "9", "10"}},
		{`new []string{}.pushMultiple("0", 10)`, []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}},
		{`new []int{1, 2, 3}.pushMultiple(0, 10).size()`, big.NewInt(13)},
		{`new []int{1, 2, 3}.slice(0, 3).size()`, big.NewInt(3)},
		{`new []int{1, 2, 3}.slice(0, 2).size()`, big.NewInt(2)},
		{`new []int{1, 2, 3}.slice(1, 2).size()`, big.NewInt(2)},
		{`new []int{1, 2, 3}.slice(1, 100).size()`, big.NewInt(2)},
		{`new []int{1, 2, 3}.slice(2, 1).size()`, big.NewInt(1)},
		{`new []int{1, 2, 3}.slice(2, 0).size()`, big.NewInt(0)},
	}
	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case *big.Int:
			testIntegerObject(t, testEval(tt.input), expected)
		case string:
			testStringObject(t, testEval(tt.input), expected)
		case []string:
			testArrayObject(t, testEval(tt.input), expected)
		default:
			t.Errorf("Test doesn't support %T expected type", expected)
		}
	}
}
