package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileReturnStatement(stmt *ast.ReturnStatement, b *context.BlockContext) *context.BlockContext {
	if stmt.ReturnValue == nil {
		b.NewRet(nil)
		return b
	}

	value := c.compileExpression(stmt.ReturnValue, b)
	b.NewRet(value)
	return b
}
