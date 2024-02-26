package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileIfStatement(expr *ast.IfStatement, block *ir.Block) *ir.Block {
	consequenceBlock := block.Parent.NewBlock("")
	continueBlock := block.Parent.NewBlock("")
	continueBlock.NewRet(nil)

	consequenceBlock.NewBr(continueBlock)
	c.compileBlockStatement(expr.Consequence, consequenceBlock)

	condition := c.compileExpression(expr.Condition, block)

	if(expr.Alternative != nil) {
		alternativeBlock := block.Parent.NewBlock("")
		
		alternativeBlock.NewBr(continueBlock)
		c.compileBlockStatement(expr.Alternative, alternativeBlock)

		block.NewCondBr(condition, consequenceBlock, alternativeBlock)
	} else {
		block.NewCondBr(condition, consequenceBlock, continueBlock)
	}

	return continueBlock
}