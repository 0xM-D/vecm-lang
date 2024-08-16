package compiler

import (
	"os"
	"testing"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"tinygo.org/x/go-llvm"
)


func compileAndVerifyCode(code string, t *testing.T) *ir.Module {
	

	c, _ := InitializeCompiler()
	_, hasParserErrors := c.loadModule("code", code);

	if hasParserErrors {
		t.Fatalf("Expected no parser errors, got some")
	}
	
	ir, hasCompilerErrors := c.CompileModule("code")

	if hasCompilerErrors {
		c.printCompilerErrors();
		t.Fatalf("Expected no compiler errors, got some")
	}

	irModule, error := asm.ParseString("", ir)

	if error != nil {
		t.Fatalf("Generated IR is invalid")
	}

	return irModule;
}

func compileModuleForExecution(ctx llvm.Context, ir string, t *testing.T) llvm.ExecutionEngine {
    // Initialize LLVM
	llvm.InitializeAllTargets()
	llvm.InitializeAllAsmPrinters()
	// llvm.InitializeAllAsmParsers()
	// llvm.InitializeAllTargetInfos()

	// Open file on os
	file, err := os.CreateTemp("", "ir")
	if err != nil {
		t.Fatalf("Failed to create temporary file")
	}
	defer file.Close()

	// Write IR to file

	_, err = file.WriteString(ir)

	if err != nil {
		t.Fatalf("Failed to write IR to file")
	}

	// New memory buffer from file
	memoryBuffer, err := llvm.NewMemoryBufferFromFile(file.Name())

	if err != nil {
		t.Fatalf("Failed to create memory buffer from file")
	}

	module, err := ctx.ParseIR(memoryBuffer)

	if err != nil {
		t.Fatalf("Failed to parse IR")
	}

	// Create a new execution engine
	engine, err := llvm.NewExecutionEngine(module)

	if err != nil {
		t.Fatalf("Failed to create execution engine")
	}

	return engine
}

func expectFunctionExists(module *ir.Module, funcName string, paramTypes []types.Type, returnType types.Type, t *testing.T) *ir.Func {
	fn, found := findFunction(module, funcName)

	if !found {
		t.Fatalf("Function %s not found", funcName)
	}

	if len(fn.Params) != len(paramTypes) {
		t.Fatalf("Function %s has wrong number of parameters", funcName)
	}

	for i, param := range fn.Params {
		if !param.Type().Equal(paramTypes[i]) {
			t.Fatalf("Function %s has wrong parameter type", funcName)
		}
	}

	if !fn.Sig.RetType.Equal(returnType) {
		t.Fatalf("Function %s has wrong return type", funcName)
	}

	return fn
}

func expectFunctionHasNBlocks(fn *ir.Func, n int, t *testing.T) []*ir.Block {

	if len(fn.Blocks) != n {
		t.Fatalf("Function %s has wrong number of blocks, ecxpected %d, got %d", fn.Name(), n, len(fn.Blocks))
	}

	return fn.Blocks
}

func expectReturnTerminator(block *ir.Block, value value.Value, t *testing.T) {
	if block.Term == nil {
		t.Fatalf("Expected terminator in block, got nil")
	}

	if block.Term.(*ir.TermRet) == nil {
		t.Fatalf("Expected return terminator in block, got nil")
	}

	if block.Term.(*ir.TermRet).X.String() != value.String() {
		t.Fatalf("Expected return value %v, got %v", value, block.Term.(*ir.TermRet).X)
	}
}

func findFunction (module *ir.Module, funcName string) (*ir.Func, bool) {
	for _, f := range module.Funcs {
		if f.Name() == funcName {
			return f, true
		}
	}

	return nil, false
}