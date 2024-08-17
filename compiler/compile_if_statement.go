package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
)

func (c *Compiler) compileIfStatement(expr *ast.IfStatement, block *context.BlockContext) *context.BlockContext {
	condition := c.compileExpression(expr.Condition, block)
	
	var consequenceBlock *context.BlockContext;
	var alternativeBlock *context.BlockContext;
	var continueBlock *context.BlockContext;

	consequenceBlock = context.NewBlockContext(block.GetParentContext(), block.Parent.NewBlock(""))
	c.compileBlockStatement(expr.Consequence, consequenceBlock);

	if expr.Alternative != nil {
		alternativeBlock = context.NewBlockContext(block.GetParentContext(), block.Parent.NewBlock(""))
		c.compileBlockStatement(expr.Alternative, alternativeBlock)
	}

	consequenceBlockHasTerm := consequenceBlock != nil && consequenceBlock.Block.Term != nil
	alternativeBlockHasTerm := alternativeBlock != nil && alternativeBlock.Block.Term != nil

	if consequenceBlockHasTerm && alternativeBlockHasTerm {
		block.NewCondBr(condition, consequenceBlock.Block, alternativeBlock.Block);
	} else {
		continueBlock = context.NewBlockContext(block.GetParentContext(), block.Parent.NewBlock(""))
		block.NewCondBr(condition, consequenceBlock.Block, continueBlock.Block)

		if !consequenceBlockHasTerm && consequenceBlock != nil {
			consequenceBlock.NewBr(continueBlock.Block)
		}

		if !alternativeBlockHasTerm && alternativeBlock != nil {
			alternativeBlock.NewBr(continueBlock.Block)
		}
	}

	return continueBlock
}