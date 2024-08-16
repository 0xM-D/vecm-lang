package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileIfStatement(expr *ast.IfStatement, block *context.BlockContext) *ir.Block {
	consequenceBlock := context.NewBlockContext(block.GetParentContext(), block.Parent.NewBlock(""))
	continueBlock := context.NewBlockContext(block.GetParentContext(), block.Parent.NewBlock(""))
	continueBlock.NewRet(nil)

	consequenceBlock.NewBr(continueBlock.Block)
	c.compileBlockStatement(expr.Consequence, consequenceBlock)

	condition := c.compileExpression(expr.Condition, block)

	if(expr.Alternative != nil) {
		alternativeBlock := context.NewBlockContext(block.GetParentContext(), block.Parent.NewBlock(""))
		
		alternativeBlock.NewBr(continueBlock.Block)
		c.compileBlockStatement(expr.Alternative, alternativeBlock)

		block.NewCondBr(condition, consequenceBlock.Block, alternativeBlock.Block)
	} else {
		block.NewCondBr(condition, consequenceBlock.Block, continueBlock.Block)
	}

	return continueBlock.Block
}