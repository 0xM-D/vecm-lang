package evaluator

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated, err := testEval(input)
	if err != nil {
		t.Fatal(err)
	}
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBooleanInfixExpression(t *testing.T) {
	tests := []struct {
		input string

		expected bool
	}{
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"true && true", true},
		{"false && false", false},
		{"false && true", false},
		{"true || true", true},
		{"false || false", false},
		{"false || true", true},
		{"(1 < 2) && (2 < 3)", true},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestFloatInfixExpression(t *testing.T) {

	tests := []struct {
		input    string
		expected interface{}
	}{
		{"5.0", 5.0},
		{"-6.0", -6.0},
		{"1.1f", float32(1.1)},
		{"-5f", float32(-5.0)},
		{"-10.22233344f", -float32(10.22233344)},
		{"2.0 + 2", 4.0},
		{"2.0f * 3", float32(6.0)},
		{"5.0 + 5.0f + .5 + 5 - 10", 5.5},
		{"2.0 * 2f * 2.0f * 2 * 2", 32.0},
		{"-50 + 100 + -50", big.NewInt(0)},
		{"5f * 2 + 10f", float32(20)},
		{"5 + 2f * 10", float32(25)},
		{"20f + 2.0 * -10f", 0.0},
		{"51 / 2 * 2f + 10f", float32(60)},
		{"2 * (5f + 10.0)", 30.0},
		{"3f * 3f * 3f + 10", float32(37.0)},
		{"3 * (3f * 3) + 10", float32(37.0)},
		{"(5.0f + 10 * 2 + 11f / 3) * 2 + -10", float32(47.333333333333336)},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testNumber(t, evaluated, tt.expected)
	}
}
