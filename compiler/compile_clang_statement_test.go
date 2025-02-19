package compiler_test

import (
	"testing"

	llvm "tinygo.org/x/go-llvm"
)

func TestCBlock(t *testing.T) {
	code := `CLang {
				int mult2(int x) {
					return x * 2;
				}
			}

			fn mult2(x: int32)->int32;

			fn main() -> int32 {
				return mult2(5);
			}
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

	if result.Int(false) != 10 {
		t.Fatalf("Expected 10, got %d", result.Int(false))
	}
}

func TestCStdLib(t *testing.T) {
	code := `CLang {
				#include <math.h>

				float sinDegrees(float x) {
					float pi = 3.14159265359;
					return sin(x * pi / 180);
				}
			}

			fn sinDegrees(x: float32)->float32;
			`

	module := compileAndVerifyCode(code, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module, t)

	// Find the function
	executableFn := executionEngine.FindFunction("sinDegrees")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		input  float32
		output float32
	}{
		{0, 0},
		{90, 1},
		{180, 0},
		{270, -1},
		{360, 0},
		{45, 0.70710678118},
	}

	for _, test := range tests {
		args := []llvm.GenericValue{
			llvm.NewGenericValueFromFloat(ctx.FloatType(), float64(test.input)),
		}
		// Execute the function
		result := executionEngine.RunFunction(executableFn, args)
		delta := 0.0001
		if diff := result.Float(ctx.FloatType()) - float64(test.output); diff < -delta || diff > delta {
			t.Fatalf("Expected %f, got %f", test.output, result.Float(ctx.FloatType()))
		}
	}
}
