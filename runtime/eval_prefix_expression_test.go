package runtime

import "testing"

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testBooleanObject(t, evaluated, tt.expected)
	}
}
