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

	if inst.LLString() != "%1 = add i32 5, 5" {
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
	if inst.LLString() != "%1 = sub i32 10, 5" {
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
	if inst.LLString() != "%1 = mul i32 2, 3" {
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
	if inst.LLString() != "%1 = sdiv i32 10, 2" {
		t.Fatalf("Expected sdiv i32 10, 2, got %s", inst.LLString())
	}
}

// TestLessThanExpression tests the less than comparison of two numbers.
func TestLessThanExpression(t *testing.T) {
	code := `fn main() -> bool { return 5 < 10; }`
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
	if inst.LLString() != "%1 = icmp slt i32 5, 10" {
		t.Fatalf("Expected icmp slt i32 5, 10, got %s", inst.LLString())
	}
}

// TestGreaterThanExpression tests the greater than comparison of two numbers.
func TestGreaterThanExpression(t *testing.T) {
	code := `fn main() -> bool { return 10 > 5; }`
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
	if inst.LLString() != "%1 = icmp sgt i32 10, 5" {
		t.Fatalf("Expected icmp sgt i32 10, 5, got %s", inst.LLString())
	}
}

// TestEqualExpression tests the equal comparison of two numbers.
func TestEqualExpression(t *testing.T) {
	code := `fn main() -> bool { return 5 == 5; }`
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
	if inst.LLString() != "%1 = icmp eq i32 5, 5" {
		t.Fatalf("Expected icmp eq i32 5, 5, got %s", inst.LLString())
	}
}

// TestNotEqualExpression tests the not equal comparison of two numbers.
func TestNotEqualExpression(t *testing.T) {
	code := `fn main() -> bool { return 5 != 10; }`
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
	if inst.LLString() != "%1 = icmp ne i32 5, 10" {
		t.Fatalf("Expected icmp ne i32 5, 10, got %s", inst.LLString())
	}
}
// TestGreaterOrEqualExpression tests the greater than or equal comparison of two numbers.
func TestGreaterOrEqualExpression(t *testing.T) {
	code := `fn main() -> bool { return 10 >= 5; }`
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
	if inst.LLString() != "%1 = icmp sge i32 10, 5" {
		t.Fatalf("Expected icmp sge i32 10, 5, got %s", inst.LLString())
	}
}

// TestLessOrEqualExpression tests the less than or equal comparison of two numbers.
func TestLessOrEqualExpression(t *testing.T) {
	code := `fn main() -> bool { return 5 <= 10; }`
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
	if inst.LLString() != "%1 = icmp sle i32 5, 10" {
		t.Fatalf("Expected icmp sle i32 5, 10, got %s", inst.LLString())
	}
}

// TestBitwiseAndExpression tests the bitwise AND operation of two numbers.
func TestBitwiseAndExpression(t *testing.T) {
	code := `fn main() -> int { return 5 & 3; }`
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
	if inst.LLString() != "%1 = and i32 5, 3" {
		t.Fatalf("Expected and i32 5, 3, got %s", inst.LLString())
	}
}

// TestBitwiseOrExpression tests the bitwise OR operation of two numbers.
func TestBitwiseOrExpression(t *testing.T) {
	code := `fn main() -> int { return 5 | 3; }`
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
	if inst.LLString() != "%1 = or i32 5, 3" {
		t.Fatalf("Expected or i32 5, 3, got %s", inst.LLString())
	}
}

// TestBitwiseXorExpression tests the bitwise XOR operation of two numbers.
func TestBitwiseXorExpression(t *testing.T) {
	code := `fn main() -> int { return 5 ^ 3; }`
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
	if inst.LLString() != "%1 = xor i32 5, 3" {
		t.Fatalf("Expected xor i32 5, 3, got %s", inst.LLString())
	}
}

// TestAssignmentExpression tests the assignment of a value to a variable.
func TestAssignmentExpression(t *testing.T) {
	code := `fn main() -> int { let x = 5; return x; }`
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
	if len(block.Insts) != 2 {
		t.Fatalf("Expected 2 instructions, got %d", len(block.Insts))
	}

	inst1 := block.Insts[0]
	inst2 := block.Insts[1]

	// Assert the instructions
	if inst1.LLString() != "%1 = alloca i32" {
		t.Fatalf("Expected alloca %%1 = alloca i32, got %s", inst1.LLString())
	}

	if inst2.LLString() != "store i32 5, i32* %1" {
		t.Fatalf("Expected store i32 5, i32* %%1, got %s", inst2.LLString())
	}
}
