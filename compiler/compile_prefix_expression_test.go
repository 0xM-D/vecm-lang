package compiler

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"tinygo.org/x/go-llvm"
)

func TestMinusPrefixExpression(t *testing.T) {
	code := `
	fn negate(a: int) -> int {
		return -a;
	}
	`
	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "negate", []types.Type{types.I32}, types.I32, t)
	blocks := expectFunctionHasNBlocks(fn, 1, t)
	block := blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0].(*ir.InstMul)

	if inst == nil {
		t.Fatalf("Expected instruction to be a multiplication, got %v", block.Insts[0])
	}

	if inst.X.String() != "i32 1" {
		t.Fatalf("Expected instruction to be i32 5, got %v", inst.X)
	}

	if inst.Y.String() != "i32 -1" {
		t.Fatalf("Expected instruction to be i32 -1, got %v", inst.Y)
	}

	expectReturnTerminator(block, inst, t)

	// Create an LLVM context
    ctx := llvm.NewContext()
    defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)
	
	// Find the function
	executableFn := executionEngine.FindFunction("main")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		input int
		expected int
	}{
		{0, 0},
		{1, -1},
		{5, -5},
		{10, -10},
		{-2, 2},
		{-5, 5},
	}

	for _, test := range tests {
		result := executionEngine.RunFunction(executableFn, []llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.input), true)})
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, result.Int(false))
		}
	}

}
func TestBangPrefixExpression(t *testing.T) {
	code := `
	fn negate(a: bool) -> bool {
		return !a;
	}
	`
	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "negate", []types.Type{types.I1}, types.I1, t)
	blocks := expectFunctionHasNBlocks(fn, 1, t)
	block := blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst, ok := block.Insts[0].(*ir.InstXor)
	if !ok {
		t.Fatalf("Expected instruction to be an XOR, got %s", block.Insts[0].LLString())
	}

	if inst.X.String() != "i1 1" {
		t.Fatalf("Expected instruction to be i1 1, got %v", inst.X)
	}

	if inst.Y.String() != "i1 true" {
		t.Fatalf("Expected instruction to be i1 true, got %v", inst.Y)
	}

	expectReturnTerminator(block, inst, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)

	// Find the function
	executableFn := executionEngine.FindFunction("main")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		input    bool
		expected bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range tests {
		result := executionEngine.RunFunction(executableFn, []llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int1Type(), uint64(boolToInt(test.input)), true)})
		if result.Int(true) != uint64(boolToInt(test.expected)) {
			t.Fatalf("Expected %v, got %v", test.expected, result.Int(false))
		}
	}
}

func TestTildePrefixExpression(t *testing.T) {
	code := `
	fn invert(a: int) -> int {
		return ~a;
	}
	`
	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "invert", []types.Type{types.I32}, types.I32, t)
	blocks := expectFunctionHasNBlocks(fn, 1, t)
	block := blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0].(*ir.InstXor)

	if inst == nil {
		t.Fatalf("Expected instruction to be a XOR, got %s", block.Insts[0].LLString())
	}

	if inst.X.String() != "i32 %1" {
		t.Fatalf("Expected operand to be i32 %%1, got %s", inst.X.String())
	}

	if inst.Y.String() != "i32 1" {
		t.Fatalf("Expected operand to be i32 1, got %v", inst.Y)
	}

	expectReturnTerminator(block, inst, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module.String(), t)

	// Find the function
	executableFn := executionEngine.FindFunction("main")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		input    int
		expected int
	}{
		{0, -1},
		{1, -2},
		{-1, 0},
		{5, -6},
		{-5, 4},
	}

	for _, test := range tests {
		result := executionEngine.RunFunction(executableFn, []llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.input), true)})
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, result.Int(false))
		}
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}


