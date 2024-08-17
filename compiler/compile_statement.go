package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
)

func (c *Compiler) compileStatement(stmt ast.Statement, b *context.BlockContext) *context.BlockContext {
	switch stmt := stmt.(type) {
	case *ast.ReturnStatement:
		return c.compileReturnStatement(stmt, b)
	case *ast.IfStatement:
		return c.compileIfStatement(stmt, b)
	}

	c.newCompilerError(stmt, "Unknown statement type: %T", stmt)

	return nil
}
