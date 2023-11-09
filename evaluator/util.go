package evaluator

import (
	"fmt"
	"math"
	"math/big"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func isTruthy(obj object.Object) bool {
	if object.IsNull(obj) {
		return false
	}

	if object.IsBoolean(obj) {
		if object.UnwrapReferenceObject(obj).(*object.Boolean).Value {
			return true
		} else {
			return false
		}
	}

	return true
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func evalExpressionsArray(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}
	return result

}

func getMinimumIntegerType(number *big.Int) (object.ObjectKind, error) {
	switch {
	case number.Cmp(big.NewInt(math.MaxInt8)) == -1:
		return object.Int8Kind, nil
	case number.Cmp(big.NewInt(math.MaxUint8)) == -1:
		return object.UInt8Kind, nil
	case number.Cmp(big.NewInt(math.MaxInt16)) == -1:
		return object.Int16Kind, nil
	case number.Cmp(big.NewInt(math.MaxUint16)) == -1:
		return object.UInt16Kind, nil
	case number.Cmp(big.NewInt(math.MaxInt32)) == -1:
		return object.Int32Kind, nil
	case number.Cmp(big.NewInt(math.MaxUint32)) == -1:
		return object.UInt32Kind, nil
	case number.Cmp(big.NewInt(math.MaxInt64)) == -1:
		return object.Int64Kind, nil
	case number.Cmp(new(big.Int).SetUint64(math.MaxUint64)) == -1:
		return object.UInt64Kind, nil
	default:
		return object.ErrorKind, fmt.Errorf("integer ouside maximum range")
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
