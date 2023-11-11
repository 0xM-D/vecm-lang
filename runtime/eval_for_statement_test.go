package runtime

import "testing"

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`string a = ""; for(int8 i = 0; i < 10; i+=1) { a += "a" } a`, "aaaaaaaaaa"},
		{"let x = 55; for(int16 i = 10; i >= 0; i-=1) { x -= i } x", 0},
		{"let i = 0; for(; i < 20; i+=1) {} i", 20},
		{"let i = 5; for(; i > 0 ;) { i-=1 } i", 0},
	}
	for _, tt := range tests {
		result, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testLiteralObject(t, result, tt.expected)
	}
}
