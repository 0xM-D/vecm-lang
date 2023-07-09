package evaluator_tests

import "testing"

func TestTypedDeclarationStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"int a = 5; a;", 5},
		{"string a = \"testmmm\"; a;", "testmmm"},
		{"array a = [1, 2, 3, 4]; let b = a; b;", []string{"1", "2", "3", "4"}},
		{"bool a = true; let b = !a; b;", false},
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
