package evaluator

import (
	"fmt"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalCallExpression(node *ast.CallExpression, env *object.Environment) (object.Object, error) {
	function, err := Eval(node.Function, env)
	if err != nil {
		return nil, err
	}

	args, err := evalExpressionsArray(node.Arguments, env)
	if err != nil {
		return nil, err
	}

	return ApplyFunction(function, args)
}

func ApplyFunction(fn object.Object, args []object.Object) (object.Object, error) {
	switch {
	case object.IsFunction(fn):
		function := object.UnwrapReferenceObject(fn).(*object.Function)

		if len(function.ParameterTypes) != len(args) {
			return nil, fmt.Errorf("incorrect parameter count for %s fun. expected=%d, got=%d", function.Type().Signature(), len(function.ParameterTypes), len(args))
		}

		extendedEnv := extendFunctionEnv(function, args)
		evaluated, err := Eval(function.Body, extendedEnv)
		if err != nil {
			return nil, err
		}
		return unwrapReturnValue(evaluated), nil
	case object.IsBuiltinFunction(fn):
		function := object.UnwrapReferenceObject(fn).(object.BuiltinFunction)

		if len(function.ParameterTypes) != len(args) {
			return nil, fmt.Errorf("incorrect parameter count for %s fun. expected=%d, got=%d", function.Type().Signature(), len(function.ParameterTypes), len(args))
		}
		params := []object.Object{}
		params = append(params, function.BoundParams...)
		params = append(params, args...)
		return function.Function(params...), nil
	default:
		return nil, fmt.Errorf("object is not a function: %s", fn.Inspect())
	}
}
