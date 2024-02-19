package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func compileReturnStatement(stmt *ast.ReturnStatement, b *ir.Block) {
	value := compileExpression(stmt.ReturnValue)
	b.NewRet(value)
}