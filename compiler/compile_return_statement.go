package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileReturnStatement(stmt *ast.ReturnStatement, b *ir.Block) *ir.Block {
	if stmt.ReturnValue == nil {
		b.NewRet(nil)
		return b
	}
	
	value := c.compileExpression(stmt.ReturnValue, b)
	b.NewRet(value)
	return b
}