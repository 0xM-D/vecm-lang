package evaluator_tests

import "testing"

func TestAsignmentExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let a = 5; a = 3;", 3},
		{"a := 5 * 5; a += 20;", 45},
		{"a := 3; b := a; a += b", 6},
		{"a := 5; a -= 1", 4},
		{"a := 5; a *= 2", 10},
		{"a := 50; a /= 5", 10},
		{`a := "a"; a += "bc"`, "abc"},
	}
	for _, tt := range tests {
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, testEval(tt.input), int64(expected))
		case int64:
			testIntegerObject(t, testEval(tt.input), expected)
		case string:
			testStringObject(t, testEval(tt.input), expected)
		case bool:
			testBooleanObject(t, testEval(tt.input), expected)
		case []string:
			testArrayObject(t, testEval(tt.input), expected)
		}
	}
}
