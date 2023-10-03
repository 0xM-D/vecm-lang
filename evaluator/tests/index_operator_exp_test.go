package evaluator_tests

import "testing"

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[]int{1, 2, 3}[0]",
			1,
		},
		{
			"[]int{1, 2, 3}[1]",
			2,
		},
		{
			"[]int{1, 2, 3}[2]",
			3,
		},
		{
			"let i = 0; []int{1}[i];",
			1,
		},
		{
			"[]int{1, 2, 3}[1 + 1];",
			3,
		},
		{
			"let myArray = []int{1, 2, 3}; myArray[2];",
			3,
		},
		{
			"let myArray = []int{1, 2, 3}; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"let myArray = []int{1, 2, 3}; let i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[]int{1, 2, 3}[3]",
			nil,
		},
		{
			"[]int{1, 2, 3}[-1]",
			nil,
		},
		{
			"let a = []int{1}; a[0] = 2; a[0]",
			2,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
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
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
		{
			"let a = {1: 2}; a[1] = 3; a[1]",
			3,
		},
		{
			"let a = {1: 2}; b := 1; a[b] = b; a[1]",
			1,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
