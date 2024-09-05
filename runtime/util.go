package runtime

import (
	"errors"
	"math"
	"math/big"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
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

func (r *Runtime) evalExpressionsArray(
	exps []ast.Expression,
	env *object.Environment,
) ([]object.Object, error) {
	result := make([]object.Object, 0, len(exps))

	for _, e := range exps {
		evaluated, err := r.Eval(e, env)
		if err != nil {
			return nil, err
		}

		result = append(result, evaluated)
	}
	return result, nil
}

func getMinimumIntegerType(number *big.Int) (object.Kind, error) {
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
		return object.AnyKind, errors.New("integer ouside maximum range")
	}
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		// TODO: Handle error
		//nolint:errcheck // I noticed this error, but now is not the time to fix it
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
