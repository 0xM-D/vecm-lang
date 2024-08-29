package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileExpression(expr ast.Expression, b *context.BlockContext) value.Value {
	switch expr := expr.(type) {
	case *ast.IntegerLiteral:
		return constant.NewInt(types.I32, expr.Value.Int64())
	case *ast.Float32Literal:
		return constant.NewFloat(types.Float, float64(expr.Value))
	case *ast.Float64Literal:
		return constant.NewFloat(types.Double, float64(expr.Value))
	case *ast.BooleanLiteral:
		return nativeBoolToLLVMBool(expr.Value)
	case *ast.Identifier:
		return c.compileIdentifierRValue(expr, b)
	case *ast.PrefixExpression:
		return c.compilePrefixExpression(expr, b)
	case *ast.InfixExpression:
		return c.compileInfixExpression(expr, b)
	case *ast.CallExpression:
		return c.compileCallExpression(expr, b)
	// case *ast.TernaryExpression:
	// 	return c.compileTernaryExpression(expr, b)

	// case *ast.TypeCastExpression:
	// 	return c.compileTypeCastExpression(expr, b)

	// For strings, arrays, hashes, etc, we will need to rething the memory model we had in our interpreter
	// case *ast.StringLiteral:
	//  	return c.compileStringLiteral(expr)
	// case *ast.ArrayLiteral:
	// 	return c.compileArrayLiteral(expr, b)
	// case *ast.HashLiteral:
	// 	return c.compileHashLiteral(expr, b)
	// case *ast.NewExpression:
	// 	return c.compileNewExpression(expr, b)
	// case *ast.AccessExpression:
	// 	return c.compileAccessExpression(expr, b)
	// case *ast.IndexExpression:
	// 	return c.compileIndexExpression(expr, b)

	// case *ast.FunctionLiteral:
	// 	return c.compileFunctionLiteral(expr, b)

	default:
		c.newCompilerError(expr, "Invalid expression node: %T", expr)
		return nil
	}
}
