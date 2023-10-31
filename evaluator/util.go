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

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	case "~":
		return evalTildePrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	if isTruthy(right) {
		return FALSE
	}
	return TRUE
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if !object.IsNumber(right) {
		return newError("Operator - not defined on type %s", right.Type().Signature())
	}

	number := right.(*object.Number)

	if object.IsInteger(number) && number.IsUnsigned() {
		return newError("Operator - not defined on unsigned integer type %s", number.Kind.Signature())
	}
	if object.IsInteger(number) && number.IsSigned() {
		return &object.Number{Value: object.Int64Bits(-number.GetInt64()), Kind: number.Kind}
	}
	if number.Kind == object.Float32Kind {
		return &object.Number{Value: uint64(math.Float32bits(-number.GetFloat32())), Kind: number.Kind}
	}
	if number.Kind == object.Float64Kind {
		return &object.Number{Value: math.Float64bits(-number.GetFloat64()), Kind: number.Kind}
	}

	return newError("Operator - not defined on number type %s", right.Type().Signature())
}

func evalTildePrefixOperatorExpression(right object.Object) object.Object {
	if !object.IsNumber(right) {
		return newError("Operator ~ not defined on type %s", right.Type().Signature())
	}

	number := right.(*object.Number)
	return &object.Number{Value: ^number.Value, Kind: number.Kind}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if object.IsError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
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

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	reference := env.GetReference(node.Value)
	if reference == nil {
		return newError("identifier not found: " + node.Value)
	}
	return reference
}

func evalExpressions(
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

func evalNewExpression(exp *ast.NewExpression,
	env *object.Environment) object.Object {

	_, isArray := exp.Type.(ast.ArrayType)
	if isArray {
		return evalNewArrayExpression(exp, env)
	}

	_, isHash := exp.Type.(ast.HashType)
	if isHash {
		return evalNewHashExpression(exp, env)
	}

	return newError("New operator not yet supported on type: %s", exp.Type.String())

}

func evalNewArrayExpression(exp *ast.NewExpression, env *object.Environment) object.Object {
	elements := evalExpressions(exp.InitializationList, env)

	arrayType := exp.Type.(ast.ArrayType)

	elementType, _ := evalType(arrayType.ElementType, env)

	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements, ArrayObjectType: object.ArrayObjectType{ElementType: elementType}}
}

func evalNewHashExpression(exp *ast.NewExpression, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for _, expr := range exp.InitializationList {
		pair, ok := expr.(*ast.PairExpression)

		if !ok {
			return newError("Found non pair element in hash initialization list")
		}

		key := object.UnwrapReferenceObject(Eval(pair.Left, env))
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type().Signature())
		}

		value := Eval(pair.Right, env)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs, HashObjectType: object.HashObjectType{KeyType: object.AnyKind, ValueType: object.AnyKind}}
}

func evalAccessExpression(left object.Object, right string, env *object.Environment) object.Object {
	repo := left.Type().Builtins()
	var member *object.BuiltinFunction
	if repo != nil {
		member = left.Type().Builtins().Get(right)
	}
	if member == nil {
		return newError("Member %s does not exist on %s", right, left.Type().Signature())
	}

	if object.IsBuiltinFunction(member) {
		return object.BuiltinFunction{BoundParams: []object.Object{left}, Function: member.Function, FunctionObjectType: member.FunctionObjectType, Name: member.Name}
	}
	return member
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case object.IsArray(left) && object.IsInteger(index):
		return evalArrayIndexExpression(left, index)
	case object.IsHash(left):
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type().Signature())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)

	idx := typeCast(index, object.Int64Kind, true).(*object.Number).GetInt64()
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return &object.ArrayElementReference{Array: arrayObject, Index: idx}
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type().Signature())
	}

	_, exists := hashObject.Pairs[key.HashKey()]
	if !exists {
		return NULL
	}

	return &object.HashElementReference{Hash: hashObject, Key: index}
}

func evalType(typeNode ast.Type, env *object.Environment) (object.ObjectType, error) {
	switch casted := typeNode.(type) {
	case ast.HashType:
		keyType, err := evalType(casted.KeyType, env)
		if err != nil {
			return nil, err
		}
		valueType, err := evalType(casted.ValueType, env)
		if err != nil {
			return nil, err
		}
		return &object.HashObjectType{KeyType: keyType, ValueType: valueType}, nil
	case ast.ArrayType:
		elementType, err := evalType(casted.ElementType, env)
		if err != nil {
			return nil, err
		}
		return &object.ArrayObjectType{ElementType: elementType}, nil
	case ast.NamedType:
		namedType, found := env.GetObjectType(casted.TokenLiteral())
		if !found {
			return nil, fmt.Errorf("Unknown type %s in: %s", casted.TokenLiteral(), typeNode.String())
		}
		return namedType, nil
	case ast.FunctionType:
		parameterTypes := []object.ObjectType{}
		returnType, err := evalType(casted.ReturnType, env)
		if err != nil {
			return nil, err
		}

		for _, p := range casted.ParameterTypes {
			paramType, err := evalType(p, env)
			if err != nil {
				return nil, err
			}
			parameterTypes = append(parameterTypes, paramType)
		}
		return &object.FunctionObjectType{ParameterTypes: parameterTypes, ReturnValueType: returnType}, nil
	}

	return nil, fmt.Errorf("Unknown type: %s", typeNode.String())
}

func evalDeclarationStatement(declNode *ast.DeclarationStatement, env *object.Environment) object.Object {
	val := object.UnwrapReferenceObject(Eval(declNode.Value, env))
	var expectedType object.ObjectType

	if declNode.Type != nil {
		var err error
		expectedType, err = evalType(declNode.Type, env)

		if err != nil {
			return newError(err.Error())
		}

	}

	if object.IsError(val) {
		return val
	}

	if object.IsNumber(val) && expectedType == nil {
		expectedType = object.Int64Kind
	}

	if expectedType != nil {
		cast := typeCast(val, expectedType, EXPLICIT_CAST)
		if !object.IsError(cast) {
			val = cast
		}
	}

	if object.IsError(val) {
		return val
	}

	if expectedType != nil && !object.TypesMatch(expectedType, val.Type()) {
		return newError("Expression of type %s cannot be assigned to %s", val.Type().Signature(), expectedType.Signature())
	}

	newObject := env.Declare(declNode.Name.Value, declNode.IsConstant, val)

	if newObject == nil {
		return newError("Identifier with name %s already exists.", declNode.Name.Value)
	}

	return newObject
}

func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	functionType, err := evalType(node.Type, env)
	if err != nil {
		return newError(err.Error())
	}

	function := &object.Function{
		Parameters:         node.Parameters,
		Env:                env,
		Body:               node.Body,
		FunctionObjectType: *functionType.(*object.FunctionObjectType),
	}
	return function
}

func evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
	forEnv := object.NewEnclosedEnvironment(env)

	if node.Initialization != nil {
		initResult := Eval(node.Initialization, forEnv)
		if object.IsError(initResult) {
			return initResult
		}
	}

	for {
		if node.Condition != nil {
			conditionResult := Eval(node.Condition, forEnv)

			if object.IsError(conditionResult) {
				return conditionResult
			}

			if !isTruthy(conditionResult) {
				break
			}
		}

		if node.Body != nil {
			bodyResult := Eval(node.Body, forEnv)

			if bodyResult != nil && object.IsError(bodyResult) {
				return bodyResult
			}
		}

		if node.AfterThought != nil {
			afterThoughtResult := Eval(node.AfterThought, forEnv)

			if object.IsError(afterThoughtResult) {
				return afterThoughtResult
			}
		}

	}

	return nil
}

func evalTernaryExpression(node *ast.TernaryExpression, env *object.Environment) object.Object {
	conditionResult := Eval(node.Condition, env)
	if object.IsError(conditionResult) {
		return conditionResult
	}

	var expressionToEvaluate ast.Node
	if isTruthy(conditionResult) {
		expressionToEvaluate = node.ValueIfTrue
	} else {
		expressionToEvaluate = node.ValueIfFalse
	}

	result := Eval(expressionToEvaluate, env)
	if object.IsError(result) {
		return result
	}

	return result
}

func evalIntegerLiteral(node *ast.IntegerLiteral, env *object.Environment) object.Object {
	kind, err := getMinimumIntegerType(&node.Value)
	if err != nil {
		return newError(err.Error())
	}

	if !object.IS_SIGNED[kind] {
		return &object.Number{Value: node.Value.Uint64(), Kind: object.UInt64Kind}
	}

	return &object.Number{Value: object.Int64Bits(node.Value.Int64()), Kind: object.Int64Kind}
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
		return object.ErrorKind, fmt.Errorf("Integer ouside maximum range")
	}
}

func evalExplicitTypeCast(node *ast.TypeCastExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if object.IsError(left) {
		return left
	}

	castToType, error := evalType(node.Type, env)
	if error != nil {
		return newError(error.Error())
	}

	casted := typeCast(left, castToType, EXPLICIT_CAST)
	if object.IsError(casted) {
		return casted
	}

	return casted
}
