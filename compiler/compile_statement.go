package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
)

func (c *Compiler) compileStatement(stmt ast.Statement, b *context.BlockContext) *context.BlockContext {
	switch stmt := stmt.(type) {
	case *ast.ReturnStatement:
		return c.compileReturnStatement(stmt, b)
	case *ast.IfStatement:
		return c.compileIfStatement(stmt, b)
	case *ast.BlockStatement:
		newBlock := context.NewBlockContext(b.GetParentContext(), b.GetParentFunctionContext().NewBlock(""))
		return c.compileBlockStatement(stmt, newBlock)
	case *ast.ExpressionStatement:
		c.compileExpression(stmt.Expression, b)
		return b
	case *ast.LetStatement:
		c.compileLetStatement(stmt, b)
		return b
	// case *ast.FunctionStatement: // OH BOY I AM NOT WRITING THIS TODAY
	// 	c.compileFunctionStatement(stmt, b)
	// 	return b
	case *ast.ForStatement:
		return c.compileForStatement(stmt, b)
	// case *ast.AssignmentDeclarationStatement:
	// 	return c.compileAssignmentDeclarationStatement(stmt, b)
	// case *ast.AssignmentStatement:
	// 	return c.compileAssignmentStatement(stmt, b)
	// case *ast.DeclarationStatement:
	// 	return c.compileDeclarationStatement(stmt, b)
	// case *ast.VariableUpdateStatement:
	// 	return c.compileVariableUpdateStatement(stmt, b)
	// case *ast.TypedDeclarationStatement:
	// 	return c.compileTypedDeclarationStatement(stmt, b)
	default:
		c.newCompilerError(stmt, "Unknown statement type: %T", stmt)
		return b
	}
}
