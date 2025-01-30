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

	ctx.LinkedModulesIR = append(ctx.LinkedModulesIR, output.String())
}
