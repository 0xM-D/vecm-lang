package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileInfixExpression(infixExpr *ast.InfixExpression, b *context.BlockContext) value.Value {
	if infixExpr.Operator == "=" ||
		infixExpr.Operator == "+=" ||
		infixExpr.Operator == "-=" ||
		infixExpr.Operator == "*=" ||
		infixExpr.Operator == "/=" {
		return c.compileAssignmentInfixExpression(infixExpr, b)
	}

	left := c.compileExpression(infixExpr.Left, b)
	right := c.compileExpression(infixExpr.Right, b)

	if left == nil || right == nil {
		return nil
	}

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
	default:
		c.newCompilerError(infixExpr, "Unknown operator: %s", infixExpr.Operator)
		return nil
	}
}

func (c *Compiler) compileAssignmentInfixExpression(infixExpr *ast.InfixExpression, b *context.BlockContext) value.Value {
	left, leftType := c.compileLValue(infixExpr.Left, b)
	right := c.compileExpression(infixExpr.Right, b)

	if left == nil || right == nil {
		return nil
	}

	switch infixExpr.Operator {
	case "=":
		b.Block.NewStore(right, left)
	case "+=":
		leftValue := b.Block.NewLoad(leftType, left)
		sum := b.Block.NewAdd(leftValue, right)
		b.Block.NewStore(sum, left)
	case "-=":
		leftValue := b.Block.NewLoad(leftType, left)
		sub := b.Block.NewSub(leftValue, right)
		b.Block.NewStore(sub, left)
	case "*=":
		leftValue := b.Block.NewLoad(leftType, left)
		mul := b.Block.NewMul(leftValue, right)
		b.Block.NewStore(mul, left)
	case "/=":
		leftValue := b.Block.NewLoad(leftType, left)
		div := b.Block.NewSDiv(leftValue, right)
		b.Block.NewStore(div, left)
	default:
		c.newCompilerError(infixExpr, "Unknown operator: %s", infixExpr.Operator)
		return nil
	}

	return left
}

func (c *Compiler) compileLValue(expr ast.Expression, b *context.BlockContext) (value.Value, types.Type) {
	switch expr := expr.(type) {
	case *ast.Identifier:
		return c.compileIdentifierLValue(expr, b)
	default:
		c.newCompilerError(expr, "Invalid lvalue: %T", expr)
		return nil, nil
	}
}
