package runtime

import (
	"testing"

	"github.com/DustTheory/interpreter/object"
)

func TestFunctionLiteral(t *testing.T) {
	input := "fn(x:int64)->int64 { x + 2; };"

	evaluated, err := testEval(input)
	if err != nil {
		t.Fatal(err)
	}

	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}
