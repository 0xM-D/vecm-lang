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
			// Create a new unreachable block

			if i == len(blockStatement.Statements)-1 {
				// If this is the last statement in the block, we don't need to create a new block
				break
			}

			parentFunctionContext, err := currentBlockContext.GetParentFunctionContext()
			if err != nil {
				c.newCompilerError(blockStatement, "%e", err)
				return nil
			}

			newBlock := parentFunctionContext.NewBlock("")

			parentContext, err := currentBlockContext.GetParentContext()
			if err != nil {
				c.newCompilerError(blockStatement, "%e", err)
				return nil
			}

			currentBlockContext = context.NewBlockContext(parentContext, newBlock)
		} else {
			currentBlockContext = nextBlockContext
		}
	}

	return currentBlockContext
}
