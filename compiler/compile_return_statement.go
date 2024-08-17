package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
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