package evaluator_tests

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/object"
)

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
	new map{string -> int }{
	"one": 10 - 9,
	two: 1 + 1,
	"thr" + "ee": 6 / 2
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, big.NewInt(expectedValue))
	}
}
