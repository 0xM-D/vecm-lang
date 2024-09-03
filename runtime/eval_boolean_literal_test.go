package runtime_test

import "testing"

func TestBooleanTrueLiteral(t *testing.T) {
	evaluated, err := testEval("true")
	if err != nil {
		t.Fatal(err)
	}
	testBooleanObject(t, evaluated, true)
}

func TestBooleanFalseLiteral(t *testing.T) {
	evaluated, err := testEval("false")
	if err != nil {
		t.Fatal(err)
	}
	testBooleanObject(t, evaluated, false)
}
