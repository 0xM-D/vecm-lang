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
		{"let arr = [1, 2, 3]; arr.push(4).size()", 4},
		{"[].push(1).size();", 1},
		{"[1, 2, 3].delete(1, 5).size();", 1},
		{`["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"].delete(1, 3);`, []string{"1", "5", "6", "7", "8", "9", "10"}},
		{`[].pushMultiple("0", 10)`, []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}},
		{`[1, 2, 3].pushMultiple(0, 10).size()`, 13},
	}
	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, testEval(tt.input), int64(expected))
		case string:
			testStringObject(t, testEval(tt.input), expected)
		case []string:
			testArrayObject(t, testEval(tt.input), expected)
		default:
			t.Errorf("Test doesn't support %T expected type", expected)
		}
	}
}
