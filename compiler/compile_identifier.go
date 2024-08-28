package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileIdentifierLValue(identifier *ast.Identifier, ctx *context.BlockContext) (value.Value, types.Type) {
	ident, ok := ctx.LookUpIdentifier(identifier.Value)
	
	if !ok {
		c.newCompilerError(identifier, "Undefined identifier: %s", identifier.Value)
		return nil, nil
	}

	return ident.GetAddress(), ident.GetType()
}

func (c *Compiler) compileIdentifierRValue(identifier *ast.Identifier, ctx *context.BlockContext) value.Value {
	ident, ok := ctx.LookUpIdentifier(identifier.Value)

	if !ok {
		c.newCompilerError(identifier, "Undefined identifier: %s", identifier.Value)
		return nil
	}

	switch ident.(type) {
		case *context.LocalVariable:
			return ctx.Block.NewLoad(ident.GetType(), ident.GetAddress());
		case *context.FunctionParamVariable:
			return ident.GetAddress()
		default:
			c.newCompilerError(identifier, "Invalid identifier type: %T", ident)
		return nil
	}
	
}