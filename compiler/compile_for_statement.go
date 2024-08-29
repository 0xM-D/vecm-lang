package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileForStatement(forStatement *ast.ForStatement, block *context.BlockContext) *context.BlockContext {
	functionContext := block.GetParentFunctionContext()

	// Create a new block for the for loop
	forBlock := context.NewBlockContext(block, functionContext.NewBlock(""))
	continueBlock := context.NewBlockContext(block, functionContext.NewBlock(""))

	// Jump into for block
	block.NewBr(forBlock.Block)

	// Compile the initialization statement
	if forStatement.Initialization != nil {
		c.compileStatement(forStatement.Initialization, forBlock)
	}

	// Create a new block for the loop body
	conditionBlock := context.NewBlockContext(forBlock, functionContext.NewBlock(""))
	bodyBlock := context.NewBlockContext(forBlock, functionContext.NewBlock(""))

	forBlock.NewBr(conditionBlock.Block)

	if forStatement.Condition == nil {
		conditionBlock.NewBr(bodyBlock.Block)
	} else {
		expressionStatement, ok := forStatement.Condition.(*ast.ExpressionStatement)

		if !ok {
			c.newCompilerError(forStatement.Condition, "Invalid for loop condition")
			return nil
		}

		condition := c.compileExpression(expressionStatement.Expression, conditionBlock)
		conditionBlock.NewCondBr(condition, bodyBlock.Block, continueBlock.Block)
	}

	// Compile the loop body
	c.compileBlockStatement(forStatement.Body, bodyBlock)

	// Compile the update statement
	if forStatement.AfterThought != nil {
		c.compileStatement(forStatement.AfterThought, bodyBlock)
	}

	// Jump back to the loop condition
	bodyBlock.NewBr(conditionBlock.Block)

	return continueBlock
}
