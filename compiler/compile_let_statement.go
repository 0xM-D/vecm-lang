package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/DustTheory/interpreter/util"
)

func (c *Compiler) compileLetStatement(stmt *ast.LetStatement, b *context.BlockContext) {
	value := c.compileExpression(stmt.Value, b)
	t := value.Type()

	if stmt.Type != nil {
		decoratorType, err := util.GetLLVMType(stmt.Type)
		if err != nil {
			c.newCompilerError(stmt, "%e", err)
			return
		}
		if !t.Equal(decoratorType) {
			c.newCompilerError(stmt, "Type mismatch: %s != %s", t, decoratorType)
			return
		}

		t = decoratorType
	}

	variable := b.DeclareLocalVariable(stmt.Name.Value, t)

	b.Block.NewStore(value, variable)
}
