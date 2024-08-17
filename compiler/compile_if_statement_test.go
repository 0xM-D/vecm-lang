package compiler

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"tinygo.org/x/go-llvm"
)

func TestIfStatement(t *testing.T) {
	code := `
	fn main() -> void {
		if 1 {
			return;
		}
		return;
	}
	`

	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "main", []types.Type{}, types.Void, t)

	blocks := expectFunctionHasNBlocks(fn, 3, t)

	if blocks[0].Term == nil {
		t.Fatalf("Expected terminator in block 0, got nil")
	}

	if blocks[0].Term.(*ir.TermCondBr) == nil {
		t.Fatalf("Expected br terminator in block 1, got nil")
	}

	if blocks[1].Term == nil {
		t.Fatalf("Expected terminator in block 0, got nil")
	}

	if blocks[1].Term.(*ir.TermRet) == nil {
		t.Fatalf("Expected ret terminator in block 1, got nil")
	}

	if blocks[2].Term == nil {
		t.Fatalf("Expected terminator in block 0, got nil")
	}

	if blocks[2].Term.(*ir.TermRet) == nil {
		t.Fatalf("Expected ret terminator in block 1, got nil")
	}


}

func TestEarlyExit(t *testing.T) {
	code := `
	fn conditional2(a: int) -> int {
		if !a {
			return 1
		}
		return 2
	}	
	`

	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "conditional2", []types.Type{types.I32}, types.I32, t)

	blocks := expectFunctionHasNBlocks(fn, 3, t)

	if blocks[0].Term == nil {
		t.Fatalf("Expected terminator in block 0, got nil")
	}

	if blocks[0].Term.(*ir.TermCondBr) == nil {
		t.Fatalf("Expected br terminator in block 1, got nil")
	}


	expectReturnTerminator(blocks[1], constant.NewInt(types.I32, 1), t)
	expectReturnTerminator(blocks[2], constant.NewInt(types.I32, 2), t)

	// Create an LLVM context
    ctx := llvm.NewContext()
    defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)
	
	// Find the function
	executableFn := executionEngine.FindFunction("conditional2")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		input int
		expected int
	}{
		{0, 1},
		{1, 2},
	}

	for _, test := range tests {
		args := []llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.input), true)}
		result := executionEngine.RunFunction(executableFn, args)
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, result.Int(true))
		}
	}

}

func TestIfElseStatement(t *testing.T) {
	code := `
	fn max(a: int, b: int) -> int {
		if a > b {
			return a;
		} else {
			return b;
		}
	}
	`

	module := compileAndVerifyCode(code, t)

	expectFunctionExists(module, "max", []types.Type{types.I32, types.I32}, types.I32, t)
	
	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)
	
	// Find the function
	executableFn := executionEngine.FindFunction("max")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{0, -1, 0},
	}

	for _, test := range tests {
		args := []llvm.GenericValue{
			llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.a), true),
			llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.b), true),
		}
		result := executionEngine.RunFunction(executableFn, args)
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, result.Int(true))
		}
	}
}
func TestNestedConditionals(t *testing.T) {
	code := `
	fn nestedConditionals(a: int, b: int, c: int) -> int {
		if a > b {
			if b > c {
				return a;
			} else {
				return c;
			}
		} else {
			if a > c {
				return b;
			} else {
				return c;
			}
		}
	}
	`

	module := compileAndVerifyCode(code, t)

	expectFunctionExists(module, "nestedConditionals", []types.Type{types.I32, types.I32, types.I32}, types.I32, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)
	
	// Find the function
	executableFn := executionEngine.FindFunction("nestedConditionals")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		a        int
		b        int
		c        int
		expected int
	}{
		{1, 2, 3, 3},
		{5, 4, 3, 5},
		{0, -1, -2, 0},
	}

	for _, test := range tests {
		args := []llvm.GenericValue{
			llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.a), true),
			llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.b), true),
			llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.c), true),
		}
		result := executionEngine.RunFunction(executableFn, args)
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, result.Int(true))
		}
	}
}
