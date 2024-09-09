package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileIdentifierLValue(
	identifier *ast.Identifier,
	ctx *context.BlockContext,
) (value.Value, types.Type) {
	ident, err := ctx.LookUpIdentifier(identifier.Value)

	if err != nil {
		c.newCompilerError(identifier, "%e", err)
		return nil, nil
	}

	return ident.GetAddress(), ident.GetType()
}

func (c *Compiler) compileIdentifierRValue(identifier *ast.Identifier, ctx *context.BlockContext) value.Value {
	ident, err := ctx.LookUpIdentifier(identifier.Value)

	if err != nil {
		c.newCompilerError(identifier, "%e", err)
		return nil
	}

	switch ident.(type) {
	case *context.LocalVariable:
		return ctx.Block.NewLoad(ident.GetType(), ident.GetAddress())
	case *context.FunctionParamVariable:
		return ident.GetAddress()
	default:
		c.newCompilerError(identifier, "Invalid identifier type: %T", ident)
		return nil
	}
}
