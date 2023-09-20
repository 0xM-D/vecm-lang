package evaluator_tests

import (
	"math/big"
	"testing"
)

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			big.NewInt(1),
		},
		{
			"[1, 2, 3][1]",
			big.NewInt(2),
		},
		{
			"[1, 2, 3][2]",
			big.NewInt(3),
		},
		{
			"let i = 0; [1][i];",
			big.NewInt(1),
		},
		{
			"[1, 2, 3][1 + 1];",
			big.NewInt(3),
		},
		{
			"let myArray = [1, 2, 3]; myArray[2];",
			big.NewInt(3),
		},
		{
			"let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			big.NewInt(6),
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]",
			big.NewInt(2),
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
		{
			"let a = [1]; a[0] = 2; a[0]",
			big.NewInt(2),
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(*big.Int)
		if ok {
			testIntegerObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			big.NewInt(5),
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			big.NewInt(5),
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			big.NewInt(5),
		},
		{
			`{true: 5}[true]`,
			big.NewInt(5),
		},
		{
			`{false: 5}[false]`,
			big.NewInt(5),
		},
		{
			`{false: 5}[false]`,
			big.NewInt(5),
		},
		{
			"let a = {1: 2}; a[1] = 3; a[1]",
			big.NewInt(3),
		},
		{
			"let a = {1: 2}; b := 1; a[b] = b; a[1]",
			big.NewInt(1),
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(*big.Int)

		if ok {
			testIntegerObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}
