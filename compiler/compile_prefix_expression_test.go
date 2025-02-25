package compiler_test

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	llvm "tinygo.org/x/go-llvm"
)

func TestMinusPrefixExpression(t *testing.T) {
	code := `
	fn negate(a: int) -> int {
		return -a;
	}
	`
	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module.CoreModule, "negate", []types.Type{types.I32}, types.I32, t)
	blocks := expectFunctionHasNBlocks(fn, 1, t)
	block := blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst, isMul := block.Insts[0].(*ir.InstMul)
	if !isMul {
		t.Fatalf("Expected instruction to be a multiplication, got %v", block.Insts[0])
	}

	if inst.X.String() != "i32 %a" {
		t.Fatalf("Expected instruction to be i32 %%a, got %v", inst.X)
	}

	if inst.Y.String() != "i32 -1" {
		t.Fatalf("Expected instruction to be i32 -1, got %v", inst.Y)
	}

	expectReturnTerminator(block, inst, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module, t)

	// Find the function
	executableFn := executionEngine.FindFunction("negate")

	if executableFn.IsNil() {
		t.Fatalf("Failed to find function")
	}

	tests := []struct {
		input    int
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
		result := executionEngine.RunFunction(
			executableFn,
			[]llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.input), true)},
		)
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, int32(result.Int(false)))
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

	fn := expectFunctionExists(module.CoreModule, "negate", []types.Type{types.I1}, types.I1, t)
	blocks := expectFunctionHasNBlocks(fn, 1, t)
	block := blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	// Expect compare with false finstruction
	inst, isCmp := block.Insts[0].(*ir.InstICmp)
	if !isCmp {
		t.Fatalf("Expected instruction to be a comparison, got %v", block.Insts[0])
	}

	if inst.Pred.String() != "eq" {
		t.Fatalf("Expected instruction to be eq, got %v", inst.Pred)
	}

	if inst.X.String() != "i1 false" {
		t.Fatalf("Expected instruction to be i1 false, got %v", inst.X)
	}

	if inst.Y.String() != "i1 %a" {
		t.Fatalf("Expected instruction to be i1 %%a, got %v", inst.Y)
	}

	expectReturnTerminator(block, inst, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module, t)

	// Find the function
	executableFn := executionEngine.FindFunction("negate")

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
		result := executionEngine.RunFunction(
			executableFn,
			[]llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int1Type(), uint64(boolToInt(test.input)), true)},
		)
		if result.Int(false) != uint64(boolToInt(test.expected)) {
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

	fn := expectFunctionExists(module.CoreModule, "invert", []types.Type{types.I32}, types.I32, t)
	blocks := expectFunctionHasNBlocks(fn, 1, t)
	block := blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst, isXor := block.Insts[0].(*ir.InstXor)
	if !isXor {
		t.Fatalf("Expected instruction to be a XOR, got %s", block.Insts[0].LLString())
	}

	if inst.X.String() != "i32 %a" {
		t.Fatalf("Expected operand to be i32 %%a, got %s", inst.X.String())
	}

	if inst.Y.String() != "i32 -1" {
		t.Fatalf("Expected operand to be i32 -1, got %v", inst.Y)
	}

	expectReturnTerminator(block, inst, t)

	// Create an LLVM context
	ctx := llvm.NewContext()
	defer ctx.Dispose()

	executionEngine := compileModuleForExecution(ctx, module, t)

	// Find the function
	executableFn := executionEngine.FindFunction("invert")

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
		result := executionEngine.RunFunction(
			executableFn,
			[]llvm.GenericValue{llvm.NewGenericValueFromInt(ctx.Int32Type(), uint64(test.input), true)},
		)
		if result.Int(true) != uint64(test.expected) {
			t.Fatalf("Expected %d, got %d", test.expected, result.Int(true))
		}
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
