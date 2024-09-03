package compiler

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func TestTypedDeclarationStatementInteger(t *testing.T) {
	code := `fn main() -> void { int a = 10; int b; }`

	module := compileAndVerifyCode(code, t)

	fn := expectFunctionExists(module, "main", []types.Type{}, types.Void, t)

	blocks := expectFunctionHasNBlocks(fn, 1, t)

	// expect 3 instructions in the block
	if len(blocks[0].Insts) != 3 {
		t.Fatalf("Expected 3 instructions in the block, got %d", len(blocks[0].Insts))
	}

	if _, ok := blocks[0].Insts[0].(*ir.InstAlloca); !ok {
		t.Fatalf("Expected an alloca instruction, got %T", blocks[0].Insts[0])
	}

	if _, ok := blocks[0].Insts[1].(*ir.InstStore); !ok {
		t.Fatalf("Expected a store instruction, got %T", blocks[0].Insts[1])
	}

	if _, ok := blocks[0].Insts[2].(*ir.InstAlloca); !ok {
		t.Fatalf("Expected an alloca instruction, got %T", blocks[0].Insts[2])
	}

}
