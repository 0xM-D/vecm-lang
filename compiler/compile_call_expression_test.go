package compiler_test

import (
	"testing"

	llvm "tinygo.org/x/go-llvm"
)

func TestCallExpression(t *testing.T) {
	code := `
		fn return1() -> int { return 1; }
		fn main() -> int { return return1(); }
	`

	module := compileAndVerifyCode(code, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module, t)

	// Find the function
	executableFn := executionEngine.FindFunction("main")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	// Execute the function
	result := executionEngine.RunFunction(executableFn, nil)

	if result.Int(false) != 1 {
		t.Fatalf("Expected 1, got %d", result.Int(false))
	}
}

func TestCallExpressionWithArguments(t *testing.T) {
	code := `
		fn add(a: int, b: int) -> int { return a + b; }
		fn main() -> int { return add(1, 2); }
	`

	module := compileAndVerifyCode(code, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module, t)

	// Find the function
	executableFn := executionEngine.FindFunction("main")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	// Execute the function
	result := executionEngine.RunFunction(executableFn, nil)

	if result.Int(false) != 3 {
		t.Fatalf("Expected 3, got %d", result.Int(false))
	}
}
