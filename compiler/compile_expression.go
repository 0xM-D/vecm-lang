package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
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
			return c.compileIdentifier(expr, b)
		case *ast.PrefixExpression:
			return c.compilePrefixExpression(expr, b)
		case *ast.InfixExpression:
			return nil
			// return r.evalInfixExpression(node, env)
		default:
			c.newCompilerError(expr, "Invalid expression node: %T", expr)
			return nil
	}
}

