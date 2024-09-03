package runtime_test

import (
	"testing"

	"github.com/DustTheory/interpreter/object"
)

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

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
