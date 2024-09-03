package compiler

import (
	"testing"

	"github.com/llir/llvm/ir/types"
)

// TODO: Test these later
// c := fn(b: int) -> int { return b * 2 };
// d := new map{int -> int}{1: 2, 2: 3};
// e := new []int{1, 2, 3, 4, 5};
// f := "string value";
// fun := fn()->void {};

func TestIntegerAssignmentDeclarationStatement(t *testing.T) {
	code := `fn main() -> void { a := 10; }`

	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "main", []types.Type{}, types.Void, t)

	blocks := expectFunctionHasNBlocks(fn, 1, t)

	// expect 2 instructions
	if len(blocks[0].Insts) != 2 {
		t.Fatalf("Expected 2 instructions, got %d", len(blocks[0].Insts))
	}

	// expect alloca instruction
	if blocks[0].Insts[0].LLString() != "%1 = alloca i32" {
		t.Fatalf("Expected alloca instruction, got %s", blocks[0].Insts[0].LLString())
	}

	// expect store instruction
	if blocks[0].Insts[1].LLString() != "store i32 10, i32* %1" {
		t.Fatalf("Expected store instruction, got %s", blocks[0].Insts[1].LLString())
	}
}

func TestBooleanAssignmentDeclarationStatement(t *testing.T) {
	code := `fn main() -> void { a := true; }`

	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "main", []types.Type{}, types.Void, t)

	blocks := expectFunctionHasNBlocks(fn, 1, t)

	// expect 2 instructions
	if len(blocks[0].Insts) != 2 {
		t.Fatalf("Expected 2 instructions, got %d", len(blocks[0].Insts))
	}

	// expect alloca instruction
	if blocks[0].Insts[0].LLString() != "%1 = alloca i1" {
		t.Fatalf("Expected alloca instruction, got %s", blocks[0].Insts[0].LLString())
	}

	// expect store instruction
	if blocks[0].Insts[1].LLString() != "store i1 true, i1* %1" {
		t.Fatalf("Expected store instruction, got %s", blocks[0].Insts[1].LLString())
	}
}
