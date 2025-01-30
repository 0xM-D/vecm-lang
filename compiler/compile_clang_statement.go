package compiler

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/llir/llvm/asm"
)

func (c *Compiler) compileCLangStatement(node *ast.CLangStatement, ctx *context.GlobalContext) {
	// Use clang to compile c code to llvm ir, then add it to the context

	cmd := exec.Command("clang", "-x", "c", "-emit-llvm", "-S", "-o", "-", "-")

	cmd.Stdin = bytes.NewReader([]byte(node.CLangCode))

	var output bytes.Buffer
	cmd.Stdout = &output

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		c.newCompilerError(node, "error running clang for LLVM IR: %v\n%s", err, stderr.String())
		return
	}

	err = cmd.Wait()
	if err != nil {
		c.newCompilerError(node, "error running clang for LLVM IR: %v\n%s", err, stderr.String())
		return
	}

	// Add the llvm ir to the context

	clangSnippetName := fmt.Sprintf("clang_snippet_%d-%d", node.Token.Linen, node.Token.Coln)

	clangSnippetModule, err := asm.ParseString(clangSnippetName, output.String())

	if err != nil {
		c.newCompilerError(node, "error parsing clang output: %v", err)
		return
	}

	ctx.LinkedModules = append(ctx.LinkedModules, clangSnippetModule)
}
