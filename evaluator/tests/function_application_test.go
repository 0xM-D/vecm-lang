package evaluator_tests

import "testing"

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x: int) -> int { x; }; identity(5);", 5},
		{"let identity = fn(x: int) -> int { return x; }; identity(5);", 5},
		{"let double = fn(x: int) -> int { x * 2; }; double(5);", 10},
		{"let add = fn(x: int, y: int) -> int { x + y; }; add(5, 5);", 10},
		{"let add = fn(x: int, y: int) -> int { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x: int)-> int { x; }(5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
