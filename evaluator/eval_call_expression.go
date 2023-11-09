package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if object.IsError(function) {
		return function
	}

	args := evalExpressionsArray(node.Arguments, env)
	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

	return applyFunction(function, args)
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
