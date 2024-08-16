package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileIdentifier(identifier *ast.Identifier, ctx context.Context) value.Value {
	ident, ok := ctx.LookUpIdentifier(identifier.Value)
	
	if !ok {
		c.newCompilerError(identifier, "Undefined identifier: %s", identifier.Value)
		return nil
	}

	return ident.GetValue()
}