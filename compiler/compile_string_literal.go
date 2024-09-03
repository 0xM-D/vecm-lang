package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// For now, string literals can be treated as constants
// This will change when we implement a different memory model
// This is untested code.
func (c *Compiler) compileStringLiteral(stringLiteral *ast.StringLiteral) value.Value {
	return constant.NewCharArrayFromString(stringLiteral.Value)
}
