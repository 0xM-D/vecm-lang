package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalCallExpression(node *ast.CallExpression, env *object.Environment) (object.Object, error) {
	function, err := r.Eval(node.Function, env)
	if err != nil {
		return nil, err
	}

	args, err := r.evalExpressionsArray(node.Arguments, env)
	if err != nil {
		return nil, err
	}

	return r.ApplyFunction(function, args)
}

func (r *Runtime) ApplyFunction(fn object.Object, args []object.Object) (object.Object, error) {
	switch {
	case object.IsFunction(fn):
		function, ok := object.UnwrapReferenceObject(fn).(*object.Function)
		if !ok {
			return nil, fmt.Errorf("object is not a function: %s", fn.Inspect())
		}

		if len(function.ParameterTypes) != len(args) {
			return nil, fmt.Errorf("incorrect parameter count for %s fun. expected=%d, got=%d",
				function.Type().Signature(), len(function.ParameterTypes), len(args))
		}

		if len(function.ParameterTypes) != len(args) {
			return nil, fmt.Errorf("incorrect parameter count for %s fun. expected=%d, got=%d",
				function.Type().Signature(), len(function.ParameterTypes), len(args))
		}

		extendedEnv, err := passParameterValuesToFunction(function, args)
		if err != nil {
			return nil, err
		}
		evaluated, err := r.Eval(function.Body, extendedEnv)
		if err != nil {
			return nil, err
		}
		return unwrapReturnValue(evaluated), nil
	case object.IsBuiltinFunction(fn):
		function, ok := object.UnwrapReferenceObject(fn).(*object.BuiltinFunction)
		if !ok {
			return nil, fmt.Errorf("object is not a builtin function: %s", fn.Inspect())
		}

		if len(function.ParameterTypes) != len(args) {
			return nil, fmt.Errorf("incorrect parameter count for %s fun. expected=%d, got=%d",
				function.Type().Signature(), len(function.ParameterTypes), len(args))
		}
		params := []object.Object{}
		params = append(params, function.BoundParams...)
		params = append(params, args...)
		return function.Function(params...), nil
	default:
		return nil, fmt.Errorf("object is not a function: %s", fn.Inspect())
	}
}
