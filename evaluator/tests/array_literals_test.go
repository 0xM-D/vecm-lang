package evaluator_tests

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestArrayLiterals(t *testing.T) {
	input := "[]int{1, 2 * 2, 3 + 3}"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)

	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], big.NewInt(1))
	testIntegerObject(t, result.Elements[1], big.NewInt(4))
	testIntegerObject(t, result.Elements[2], big.NewInt(6))

}
