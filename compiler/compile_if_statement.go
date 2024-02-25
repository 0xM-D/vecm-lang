package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileIfStatement(expr *ast.IfStatement, block *ir.Block) {
	consequenceBlock := block.Parent.NewBlock("if.consequence")
	alternativeBlock := block.Parent.NewBlock("if.alternative")

	c.compileBlock(expr.Consequence, consequenceBlock)
	c.compileBlock(expr.Alternative, alternativeBlock)

	condition := c.compileExpression(expr.Condition, block)
	block.NewCondBr(condition, consequenceBlock, alternativeBlock)

}