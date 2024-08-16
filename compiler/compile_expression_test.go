package compiler

import "testing"

// TestAddExpression tests the addition of two numbers.
func TestAddExpression(t *testing.T) {
	code := `fn main() -> int { return 5 + 5; }`
	module := compileAndVerifyCode(code, t)

	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	if len(fn.Blocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(fn.Blocks))
	}

	block := fn.Blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0]

	if inst.LLString() != "add i32 5, 5" {
		t.Fatalf("Expected add i32 5, 5, got %s", inst.LLString())
	}
}

// TestSubtractExpression tests the subtraction of two numbers.
func TestSubtractExpression(t *testing.T) {
	code := `fn main() -> int { return 10 - 5; }`
	module := compileAndVerifyCode(code, t)

	// Assert the number of functions
	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	// Assert the number of blocks
	if len(fn.Blocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(fn.Blocks))
	}

	block := fn.Blocks[0]

	// Assert the number of instructions
	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0]

	// Assert the instruction
	if inst.LLString() != "sub i32 10, 5" {
		t.Fatalf("Expected sub i32 10, 5, got %s", inst.LLString())
	}
}

// TestMultiplyExpression tests the multiplication of two numbers.
func TestMultiplyExpression(t *testing.T) {
	code := `fn main() -> int { return 2 * 3; }`
	module := compileAndVerifyCode(code, t)

	// Assert the number of functions
	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	// Assert the number of blocks
	if len(fn.Blocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(fn.Blocks))
	}

	block := fn.Blocks[0]

	// Assert the number of instructions
	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0]

	// Assert the instruction
	if inst.LLString() != "mul i32 2, 3" {
		t.Fatalf("Expected mul i32 2, 3, got %s", inst.LLString())
	}
}

// TestDivideExpression tests the division of two numbers.
func TestDivideExpression(t *testing.T) {
	code := `fn main() -> int { return 10 / 2; }`
	module := compileAndVerifyCode(code, t)

	// Assert the number of functions
	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	// Assert the number of blocks
	if len(fn.Blocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(fn.Blocks))
	}

	block := fn.Blocks[0]

	// Assert the number of instructions
	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0]

	// Assert the instruction
	if inst.LLString() != "sdiv i32 10, 2" {
		t.Fatalf("Expected sdiv i32 10, 2, got %s", inst.LLString())
	}
}
