package runtime

import (
	"math"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func (r *Runtime) Eval(node ast.Node, env *object.Environment) (object.Object, error) {
	switch node := node.(type) {
	case *ast.Program:
		return r.evalProgram(node, env)
	case *ast.ExpressionStatement:
		return r.Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return r.evalIntegerLiteral(node, env)
	case *ast.Float32Literal:
		return &object.Number{Value: uint64(math.Float32bits(node.Value)), Kind: object.Float32Kind}, nil
	case *ast.Float64Literal:
		return &object.Number{Value: math.Float64bits(node.Value), Kind: object.Float64Kind}, nil
	case *ast.BooleanLiteral:
		return r.evalBooleanLiteral(node)
	case *ast.PrefixExpression:
		return r.evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		return r.evalInfixExpression(node, env)
	case *ast.BlockStatement:
		return r.evalBlockStatement(node, env)
	case *ast.IfExpression:
		return r.evalIfExpression(node, env)
	case *ast.ReturnStatement:
		return r.evalReturnStatement(node, env)
	case *ast.LetStatement:
		return r.evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.Identifier:
		return r.evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		return r.evalFunctionLiteral(node, env)
	case *ast.CallExpression:
		return r.evalCallExpression(node, env)
	case *ast.StringLiteral:
		return r.evalStringLiteral(node)
	case *ast.IndexExpression:
		return r.evalIndexExpression(node, env)
	case *ast.AccessExpression:
		return r.evalAccessExpression(node, env)
	case *ast.TypedDeclarationStatement:
		return r.evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.AssignmentDeclarationStatement:
		return r.evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.ForStatement:
		return r.evalForStatement(node, env)
	case *ast.TernaryExpression:
		return r.evalTernaryExpression(node, env)
	case *ast.TypeCastExpression:
		return r.evalExplicitTypeCast(node, env)
	case *ast.NewExpression:
		return r.evalNewExpression(node, env)
	case *ast.ImportStatement:
		return r.evalImportStatement(node, env)
	case *ast.ExportStatement:
		return r.evalExportStatement(node, env)
	case *ast.FunctionDeclarationStatement:
		return r.evalFunctionDeclarationStatement(node, env)
	}

	return nil, nil
}

func (r *Runtime) evalProgram(program *ast.Program, env *object.Environment) (object.Object, error) {
	var result object.Object
	var err error
	for _, statement := range program.Statements {
		result, err = r.Eval(statement, env)

		if err != nil {
			return nil, err
		}
		if result != nil && object.IsReturnValue(result) {
			return unwrapReturnValue(result), nil
		}
	}
	return result, nil
}
