package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func compileStatement(stmt ast.Statement, b *ir.Block) {
	switch stmt := stmt.(type) {
	case *ast.ReturnStatement:
		compileReturnStatement(stmt, b)
	}
}
