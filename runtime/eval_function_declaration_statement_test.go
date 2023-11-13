package runtime

import (
	"math/big"
	"testing"
)

func TestFunctionDeclarationStatement(t *testing.T) {
	input := `
	fn functionName(x: int64)->int64 {
		return x * 2;
	}
	functionName(50)
	`

	evaluated, err := testEval(input)
	if err != nil {
		t.Fatal(err)
	}

	testIntegerObject(t, evaluated, big.NewInt(100))
}
