package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileBlockStatement(blockStatement *ast.BlockStatement, entry *ir.Block) *ir.Block {
	currentBlock := entry
	for _, stmt := range blockStatement.Statements {
		currentBlock = c.compileStatement(stmt, currentBlock)
	}
	return currentBlock
}
