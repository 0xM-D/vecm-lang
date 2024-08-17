package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileInfixExpression(expr ast.Expression, b *context.BlockContext) value.Value {
	infixExpr := expr.(*ast.InfixExpression)

	left := c.compileExpression(infixExpr.Left, b)
	right := c.compileExpression(infixExpr.Right, b)

	switch infixExpr.Operator {
	case "+":
		return b.Block.NewAdd(left, right)
	case "-":
		return b.Block.NewSub(left, right)
	case "*":
		return b.Block.NewMul(left, right)
	case "/":
		return b.Block.NewSDiv(left, right)
	case "<":
		return b.Block.NewICmp(enum.IPredSLT, left, right)
	case ">":
		return b.Block.NewICmp(enum.IPredSGT, left, right)
	case "==":
		return b.Block.NewICmp(enum.IPredEQ, left, right)
	case "!=":
		return b.Block.NewICmp(enum.IPredNE, left, right)
	case "<=":
		return b.Block.NewICmp(enum.IPredSLE, left, right)
	case ">=":
		return b.Block.NewICmp(enum.IPredSGE, left, right)
	case "&&":
		return b.Block.NewAnd(left, right)
	case "||":
		return b.Block.NewOr(left, right)
	case "&":
		return b.Block.NewAnd(left, right)
	case "|":
		return b.Block.NewOr(left, right)
	case "^":
		return b.Block.NewXor(left, right)
	case "~":
		return b.Block.NewXor(left, right)
	case "<<":
		return b.Block.NewShl(left, right)
	case ">>":
		return b.Block.NewLShr(left, right)
	case "+=":
		// add and store result in left, return reference to left
		result := b.Block.NewAdd(left, right)
		b.Block.NewStore(result, left)
		return left
	case "-=":
		// subtract and store result in left, return reference to left
		result := b.Block.NewSub(left, right)
		b.Block.NewStore(result, left)
		return left
	case "*=":
		// multiply and store result in left, return reference to left
		result := b.Block.NewMul(left, right)
		b.Block.NewStore(result, left)
		return left
	case "/=":
		// divide and store result in left, return reference to left
		result := b.Block.NewSDiv(left, right)
		b.Block.NewStore(result, left)
		return left
	case "=":
		b.Block.NewStore(right, left)
		return left	
	default:
		c.newCompilerError(infixExpr, "Unknown operator: %s", infixExpr.Operator)
		return nil
	}
}

