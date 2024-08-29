package compiler

import (
	"testing"

	llvm "tinygo.org/x/go-llvm"
)

func TestForStatementWithoutInitOrAfterThought(t *testing.T) {
	code := `fn main() -> void {
		for (;;) {}
		return;
	}`

	compileAndVerifyCode(code, t)
}

func TestForStatementWithInit(t *testing.T) {
	code := `fn main() -> void {
		for (let i = 0;;) {}
		return;
	}`

	compileAndVerifyCode(code, t)
}

func TestForStatementWithAfterThought(t *testing.T) {
	code := `fn main() -> void {
		for (;; let x = 0) {}
		return;
	}`

	compileAndVerifyCode(code, t)
}

func TestForLoopCounter(t *testing.T) {
	code := `fn counter(n: int) -> int {
		let x = 0;
		for (let i = 0; i < n; i+=1) {
			x += 1;
		}
		return x;
	}`

	module := compileAndVerifyCode(code, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)

	// Find the function
	executableFn := executionEngine.FindFunction("counter")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	// Execute the function for values of n from 0 to 10
	for i := 0; i < 10; i++ {
		args := []llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(i), false)}
		result := executionEngine.RunFunction(executableFn, args)
		if result.Int(true) != uint64(i) {
			t.Fatalf("Expected %d, got %d", i, result.Int(true))
		}
	}
}
