package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileExportStatement(stmt *ast.ExportStatement, ctx *context.GlobalContext) {
	// Don't do anything with export statements at this point, let's ignore this node
	c.compileTopLevelStatement(stmt.Statement, ctx)
}
