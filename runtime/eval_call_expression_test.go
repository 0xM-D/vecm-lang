package runtime_test

import (
	"math/big"
	"testing"
)

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x: int8) -> int8 { x; }; identity(5);", 5},
		{"let identity = fn(x: int16) -> int16 { return x; }; identity(5);", 5},
		{"let double = fn(x: int32) -> int32 { x * 2; }; double(5);", 10},
		{"let add = fn(x: int64, y: int32) -> int64 { x + y; }; add(5, 5);", 10},
		{"let add = fn(x: int64, y: int64) -> int64 { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x: uint8)-> uint8 { x; }(5)", 5},
		{"a := 3; b := 4; let add = fn(x: uint32, y: uint32) -> uint32 { x + y; }; add(a, b);", 7},
	}
	for _, tt := range tests {
		result, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testIntegerObject(t, result, big.NewInt(tt.expected))
	}
}
