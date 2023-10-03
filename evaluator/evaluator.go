package evaluator

import (
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
		return &object.Number[float32]{Value: node.Value}
	case *ast.Float64Literal:
		return &object.Number[float64]{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if object.IsError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if object.IsError(val) {
			return val
		}
		return &object.ReturnValue{Value: val, ReturnValueObjectType: object.ReturnValueObjectType{ReturnType: val.Type()}}
	case *ast.LetStatement:
		return evalDeclarationStatement(&(*node).DeclarationStatement, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if object.IsError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && object.IsError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.ArrayLiteral:
		return evalArrayLiteral(node, env)
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if object.IsError(index) {
			return index
		}
		return evalIndexExpression(object.UnwrapReferenceObject(left), object.UnwrapReferenceObject(index))
	case *ast.AccessExpression:
		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}
		right, ok := node.Right.(*ast.Identifier)
		if !ok {
			return newError("Right side of access expression is not an identifier")
		}
		return evalAccessExpression(object.UnwrapReferenceObject(left), right.Value, env)

	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
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

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			if object.IsError(result) || object.IsReturnValue(result) {
				return result
			}
		}
	}
	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch {
	case object.IsFunction(fn):
		function := object.UnwrapReferenceObject(fn).(*object.Function)

		if len(function.ParameterTypes) != len(args) {
			return newError("Incorrect parameter count for %s fun. expected=%d, got=%d", function.Type().Signature(), len(function.ParameterTypes), len(args))
		}

		extendedEnv := extendFunctionEnv(function, args)
		evaluated := Eval(function.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case object.IsBuiltinFunction(fn):
		function := object.UnwrapReferenceObject(fn).(object.BuiltinFunction)

		if len(function.ParameterTypes) != len(args) {
			return newError("Incorrect parameter count for %s fun. expected=%d, got=%d", function.Type().Signature(), len(function.ParameterTypes), len(args))
		}
		params := []object.Object{}
		params = append(params, function.BoundParams...)
		params = append(params, args...)
		return function.Function(params...)
	default:
		return newError("object is not a function: %s", fn.Inspect())
	}
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Declare(param.Value, false, object.UnwrapReferenceObject(args[paramIdx]))
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}
