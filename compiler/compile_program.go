package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileProgram(program *ast.Program) *context.GlobalContext {
	ctx := context.NewGlobalContext(nil)

	for _, statement := range program.Statements {
		c.compileTopLevelStatement(statement, ctx)
	}

	return ctx
}

func (c *Compiler) compileTopLevelStatement(statement ast.Statement, ctx *context.GlobalContext) {
	switch statement := statement.(type) {
	case *ast.FunctionDeclarationStatement:
		c.compileFunctionDeclaration(statement, ctx)
	case *ast.ExportStatement:
		c.compileExportStatement(statement, ctx)
	default:
		c.newCompilerError(statement, "Invalid statement on top level: %T", statement)
	}
}
