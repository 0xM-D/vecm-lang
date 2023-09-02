package evaluator_tests

import "testing"

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`string a = ""; for(int i = 0; i < 10; i+=1) { a += "a" }; a`, "aaaaaaaaaa"},
		{"int x = 55; for(int i = 10; i >= 0; i-=1) { x -= i }; x", 0},
		{"int i = 0; for(; i < 20; i+=1) {}; i", 20},
		{"int i = 5; for(; i > 0 ;) { i-=1 } i", 0},
	}
	for _, tt := range tests {
		testLiteralObject(t, testEval(tt.input), tt.expected)
	}
}
