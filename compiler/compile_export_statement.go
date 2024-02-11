package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileExportStatement(stmt *ast.ExportStatement, m *ir.Module) error {
	// Don't do anything with export statements at this point, let's ignore this node
	return c.compileTopLevelStatement(stmt.Statement, m)
}
