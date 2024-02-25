package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileStatement(stmt ast.Statement, b *ir.Block) {
	switch stmt := stmt.(type) {
	case *ast.ReturnStatement:
		c.compileReturnStatement(stmt, b)
	case *ast.IfStatement:
		c.compileIfStatement(stmt, b);
	}
}
