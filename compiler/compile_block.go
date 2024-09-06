package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileBlockStatement(
	blockStatement *ast.BlockStatement,
	block *context.BlockContext,
) *context.BlockContext {
	currentBlockContext := block

	for i, stmt := range blockStatement.Statements {
		nextBlockContext := c.compileStatement(stmt, currentBlockContext)
		if nextBlockContext == nil {
			// Anything after here will be unreachable code
			// TODO: Throw a warning
			// Create a new unreachable block

			if i == len(blockStatement.Statements)-1 {
				// If this is the last statement in the block, we don't need to create a new block
				break
			}

			newBlock := currentBlockContext.GetParentFunctionContext().NewBlock("")
			currentBlockContext = context.NewBlockContext(currentBlockContext.GetParentContext(), newBlock)
		} else {
			currentBlockContext = nextBlockContext
		}
	}

	return currentBlockContext
}
