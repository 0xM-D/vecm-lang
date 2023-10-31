package evaluator_tests

import "testing"

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"new []int{1, 2, 3}[0]",
			1,
		},
		{
			"new []int{1, 2, 3}[1]",
			2,
		},
		{
			"new []int{1, 2, 3}[2]",
			3,
		},
		{
			"let i = 0; new []int{1}[i];",
			1,
		},
		{
			"new []int{1, 2, 3}[1 + 1];",
			3,
		},
		{
			"let myArray = new []int{1, 2, 3}; myArray[2];",
			3,
		},
		{
			"let myArray = new []int{1, 2, 3}; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"let myArray = new []int{1, 2, 3}; let i = myArray[0]; myArray[i]",
			2,
		},
		{
			"new []int{1, 2, 3}[3]",
			nil,
		},
		{
			"new []int{1, 2, 3}[-1]",
			nil,
		},
		{
			"let a = new []int{1}; a[0] = 2; a[0]",
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
			`new map{string->int}{"foo": 5}["foo"]`,
			5,
		},
		{
			`new map{string->int}{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; new map{string->int}{"foo": 5}[key]`,
			5,
		},
		{
			`new map{string->int}{}["foo"]`,
			nil,
		},
		{
			`new map{int->int}{5: 5}[5]`,
			5,
		},
		{
			`new map{bool->int}{true: 5}[true]`,
			5,
		},
		{
			`new map{vool->int}{false: 5}[false]`,
			5,
		},
		{
			`new map{bool->int}{false: 5}[false]`,
			5,
		},
		{
			"let a = new map{string->int}{1: 2}; a[1] = 3; a[1]",
			3,
		},
		{
			"let a = new map{string->int}{1: 2}; b := 1; a[b] = b; a[1]",
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
