package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileReturnStatement(stmt *ast.ReturnStatement, b *context.BlockContext) *ir.Block {
	if stmt.ReturnValue == nil {
		b.NewRet(nil)
		return b.Block
	}
	
	value := c.compileExpression(stmt.ReturnValue, b)
	b.NewRet(value)
	return b.Block
}