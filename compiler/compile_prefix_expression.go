package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compilePrefixExpression(expr *ast.PrefixExpression, b *ir.Block) value.Value {
	switch expr.Operator {
	case "!":
		return c.compileBangPrefixExpression(expr, b)
	// case "-":
	// 	return r.evalMinusPrefixOperatorExpression(right)
	// case "~":
	// 	return r.evalTildePrefixOperatorExpression(right)
	default:
		c.newCompilerError(expr, "unknown operator: %s", expr.Operator)
		return nil
	}
}

func (c *Compiler) compileBangPrefixExpression(expr *ast.PrefixExpression, b *ir.Block) value.Value {
	right := c.compileExpression(expr.Right, b)

	if(!types.IsInt(right.Type())) {
		c.newCompilerError(expr, "! operator not defined for type: %s", right.Type().LLString())
		return nil
	}

	return ir.NewICmp(enum.IPredEQ, constant.NewInt(right.Type().(*types.IntType), 0), right)
}