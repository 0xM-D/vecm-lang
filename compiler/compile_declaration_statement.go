package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/DustTheory/interpreter/util"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

func (c *Compiler) compileDeclarationStatement(
	stmt *ast.DeclarationStatement,
	b *context.BlockContext,
) *context.BlockContext {
	var t types.Type
	var value value.Value

	if stmt.Value == nil && stmt.Type == nil {
		c.newCompilerError(stmt, "Type needs to be specified or inferred")
		return nil
	}

	// If the type is specified, get the type
	decoratorType, err := getDecoratorType(stmt)
	if err != nil {
		c.newCompilerError(stmt, "%e", err)
		return nil
	}

	// If a statement has set value, compile the expression, and check if the type matches the decorator
	if stmt.Value != nil {
		value = c.compileExpression(stmt.Value, b)
		t = value.Type()

		if decoratorType != nil {
			if !t.Equal(decoratorType) {
				c.newCompilerError(stmt, "Type mismatch: %s != %s", t, decoratorType)
				return nil
			}
		}
	} else {
		t = decoratorType
	}

	variable := b.DeclareLocalVariable(stmt.Name.Value, t)

	if stmt.Value != nil {
		b.Block.NewStore(value, variable)
	}

	return b
}

func getDecoratorType(stmt *ast.DeclarationStatement) (types.Type, error) {
	if stmt.Type != nil {
		t, err := util.GetLLVMType(stmt.Type)
		return t, errors.Wrap(err, "Failed to get LLVM type")
	}

	//nolint:nilnil // nil is a valid return value
	return nil, nil
}
