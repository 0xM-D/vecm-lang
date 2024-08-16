package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileBlockStatement(blockStatement *ast.BlockStatement, entry *context.BlockContext) *ir.Block {
	currentBlock := entry
	for _, stmt := range blockStatement.Statements {
		nextBlock := c.compileStatement(stmt, currentBlock)
		currentBlock = context.NewBlockContext(currentBlock.GetParentContext(), nextBlock)
	}
	return currentBlock.Block
}
