package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileCallExpression(expr *ast.CallExpression, b *context.BlockContext) value.Value {
	// Evaluate parameters and their types
	var paramValues = make([]value.Value, 0, len(expr.Arguments))
	var paramTypes = make([]*ir.Param, 0, len(expr.Arguments))
	for _, param := range expr.Arguments {
		paramValue := c.compileExpression(param, b)
		paramValues = append(paramValues, paramValue)
		// Does not differentiate between const and non-const parameters
		paramTypes = append(paramTypes, &ir.Param{Typ: paramValue.Type()})
	}

	// Assume function is called by identifier
	fnName := expr.Function.(*ast.Identifier).Value

	// Generate llvm function signature from function name and argument types
	fnPtr, err := b.GetFunction(fnName, paramTypes...)
	if err != nil {
		c.newCompilerError(expr, "%e", err)
		return nil
	}

	// Generate call instruction - this doesn't seem to change the block terminator
	result := b.NewCall(fnPtr, paramValues...)

	// Return value of function call
	return result
}
