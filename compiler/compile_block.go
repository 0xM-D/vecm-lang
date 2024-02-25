package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileBlock(blockStatement *ast.BlockStatement, block *ir.Block) {
	for _, stmt := range blockStatement.Statements {
		c.compileStatement(stmt, block)
	}
}
