package evaluator

import (
	"math"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return evalIntegerLiteral(node, env)
	case *ast.Float32Literal:
		return &object.Number{Value: uint64(math.Float32bits(node.Value)), Kind: object.Float32Kind}
	case *ast.Float64Literal:
		return &object.Number{Value: math.Float64bits(node.Value), Kind: object.Float64Kind}
	case *ast.BooleanLiteral:
		return evalBooleanLiteral(node)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.LetStatement:
		return evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)
	case *ast.CallExpression:
		return evalCallExpression(node, env)
	case *ast.StringLiteral:
		return evalStringLiteral(node)
	case *ast.IndexExpression:
		return evalIndexExpression(node, env)
	case *ast.AccessExpression:
		return evalAccessExpression(node, env)
	case *ast.TypedDeclarationStatement:
		return evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.AssignmentDeclarationStatement:
		return evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.ForStatement:
		return evalForStatement(node, env)
	case *ast.TernaryExpression:
		return evalTernaryExpression(node, env)
	case *ast.TypeCastExpression:
		return evalExplicitTypeCast(node, env)
	case *ast.NewExpression:
		return evalNewExpression(node, env)
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		if result != nil {
			if object.IsError(result) {
				return result
			}
			if object.IsReturnValue(result) {
				return unwrapReturnValue(result)
			}
		}
	}
	return result
}
