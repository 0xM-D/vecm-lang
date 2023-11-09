package evaluator

import "testing"

func TestBooleanTrueLiteral(t *testing.T) {
	evaluated := testEval("true")
	testBooleanObject(t, evaluated, true)
}

func TestBooleanFalseLiteral(t *testing.T) {
	evaluated := testEval("false")
	testBooleanObject(t, evaluated, false)
}
