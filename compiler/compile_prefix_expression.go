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
	case "-":
		return c.compileMinusPrefixExpression(expr, b)
	case "~":
		return c.compileTildePrefixExpression(expr, b)
	default:
		c.newCompilerError(expr, "unknown operator: %s", expr.Operator)
		return nil
	}
}

func (c *Compiler) compileBangPrefixExpression(expr *ast.PrefixExpression, b *ir.Block) value.Value {
	right := c.compileExpression(expr.Right, b)

	if(!types.IsInt(right.Type())) {
		c.newCompilerError(expr, "operator ! not defined for type: %s", right.Type().LLString())
		return nil
	}

	return b.NewICmp(enum.IPredEQ, constant.NewInt(right.Type().(*types.IntType), 0), right)
}

func (c *Compiler) compileMinusPrefixExpression(expr *ast.PrefixExpression, b *ir.Block) value.Value {
	right := c.compileExpression(expr.Right, b)

	if(types.IsInt(right.Type())) {
		return b.NewMul(right, constant.NewInt(right.Type().(*types.IntType), -1))
	} else if(types.IsFloat(right.Type())) {
		return b.NewFMul(right, constant.NewFloat(right.Type().(*types.FloatType), -1))
	} else {
		c.newCompilerError(expr, "operator - not defined for type: %s", right.Type().LLString())
		return nil
	}
}

func (c *Compiler) compileTildePrefixExpression(expr *ast.PrefixExpression, b *ir.Block) value.Value {
	right := c.compileExpression(expr.Right, b)

	var bitmaskType *types.IntType

	if(types.IsInt(right.Type())) {
		bitmaskType = right.Type().(*types.IntType)
	} else {
		switch right.Type() {
			case types.Float:
				bitmaskType = &types.IntType{BitSize: 32}
			case types.Double:
				bitmaskType = &types.IntType{BitSize: 64}
			case types.I32Ptr:
				bitmaskType = &types.IntType{BitSize: 32}
			case types.I64Ptr:
				bitmaskType = &types.IntType{BitSize: 64}
			default:
				c.newCompilerError(expr, "operator ~ not defined for type %s", right.Type().LLString())
				return nil
		}
	}

	return b.NewXor(right, constant.NewInt(bitmaskType, ^0))
}
