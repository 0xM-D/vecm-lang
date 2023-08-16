package evaluator_tests

import "testing"

func TestTypeBuiltins(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let arr = [1, 2, 3]; arr.size();", 3},
		{"let arr = []; arr.size();", 0},
		{"[1, 2, 3, 4, 5].size();", 5},
		{`str := "abcdef"; str.length();`, 6},
		{`const string str = ""; str.length();`, 0},
		{`"bleh".length()`, 4},
		{"1.toString()", "1"},
		{"(123*456).toString()", "56088"},
	}
	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, testEval(tt.input), int64(expected))
		case string:
			testStringObject(t, testEval(tt.input), expected)
		default:
			t.Errorf("Test doesn't support %T expected type", expected)
		}
	}
}
