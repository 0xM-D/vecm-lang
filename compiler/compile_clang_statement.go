package compiler

import (
	"bytes"
	"os/exec"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileCLangStatement(node *ast.CLangStatement, ctx *context.GlobalContext) {
	// Use clang to compile c code to llvm ir, then add it to the context

	cmd := exec.Command("clang", "-x", "c", "-emit-llvm", "-S", "-o", "-", "-")

	cmd.Stdin = bytes.NewReader([]byte(node.CLangCode))

	var output bytes.Buffer
	cmd.Stdout = &output

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		c.newCompilerError(node, "error running clang for LLVM IR: %v\n%s", err, stderr.String())
		return
	}

	// clangSnippetName := fmt.Sprintf("clang_snippet_%d-%d", node.Token.Linen, node.Token.Coln)

	// clangSnippetModule, err := asm.ParseBytes(clangSnippetName, output.Bytes())

	// if err != nil {
	// 	c.newCompilerError(node, "error parsing clang output: %v", err)
	// 	return
	// }

	ctx.LinkedModules = append(ctx.LinkedModules, output.String())
}
