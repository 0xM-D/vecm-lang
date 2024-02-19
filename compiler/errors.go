package compiler

import (
	"fmt"

	"github.com/0xM-D/interpreter/ast"
)


type CompilerError struct {
	message string
	linen int
	coln int
}

func (c *Compiler) newCompilerError(node ast.Node, format string, a ...any) {
	linen := node.TokenValue().Linen
	coln := node.TokenValue().Coln

	c.Errors = append(c.Errors, CompilerError{
		message: fmt.Sprintf(format, a...),
		linen: linen,
		coln: coln,
	})
}

func (c *Compiler) hasCompilerErrors() bool {
	return len(c.Errors) != 0
}

func (c *Compiler) printCompilerErrors() {
	for _, error := range(c.Errors) {
		fmt.Print(error.message)
	}
}