package runtime_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/DustTheory/interpreter/object"
	"github.com/DustTheory/interpreter/runtime"
)

func testEval(input string) (object.Object, error) {
	r, hasErrors := runtime.NewRuntimeFromCode(input)
	if hasErrors {
		return nil, errors.New("failed to load module")
	}
	return r.Eval(r.EntryModule.Program, object.NewEnvironment())
}

func testNullObject(t *testing.T, obj object.Object) {
	if obj != runtime.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected *big.Int) bool {
	if !object.IsInteger(obj) {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	resultString := obj.Inspect()

	if resultString != expected.String() {
		t.Errorf("object has wrong value. got=%s, want=%s",
			resultString, expected.String())
		return false
	}
	return true
}

func testFloat32Object(t *testing.T, obj object.Object, expected float32) bool {
	if !object.IsFloat32(obj) {
		t.Errorf("object is not Float32. got=%T (%+v)", obj, obj)
		return false
	}

	result := object.UnwrapReferenceObject(obj).(*object.Number).GetFloat32()

	if result != expected {
		t.Errorf("object has wrong value. got=%f, want=%f",
			result, expected)
		return false
	}
	return true
}

func testFloat64Object(t *testing.T, obj object.Object, expected float64) bool {
	if !object.IsFloat64(obj) {
		t.Errorf("object is not Float64. got=%T (%+v)", obj, obj)
		return false
	}

	result := object.UnwrapReferenceObject(obj).(*object.Number).GetFloat64()

	if result != expected {
		t.Errorf("object has wrong value. got=%f, want=%f",
			result, expected)
		return false
	}
	return true
}

func testNumber(t *testing.T, obj object.Object, expected interface{}) bool {
	switch v := expected.(type) {
	case *big.Int:
		return testIntegerObject(t, obj, v)
	case float32:
		return testFloat32Object(t, obj, v)
	case float64:
		return testFloat64Object(t, obj, v)
	default:
		t.Errorf("Invalid expected type got=%T", expected)
		return false
	}
}

func testStringObject(t *testing.T, obj object.Object, expected string) {
	if !object.IsString(obj) {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
	}

	result, isString := object.UnwrapReferenceObject(obj).(*object.String)

	if !isString {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
	}
}

func testArrayObject(t *testing.T, obj object.Object, expected []interface{}) {
	if !object.IsArray(obj) {
		t.Errorf("object is not Array. got=%T (%+v)", obj, obj)
	}

	result, isArray := object.UnwrapReferenceObject(obj).(*object.Array)

	if !isArray {
		t.Errorf("object is not Array. got=%T (%+v)", obj, obj)
	}

	if len(result.Elements) != len(expected) {
		t.Errorf("Incorrect array length. expected=%d. got=%d", len(expected), len(result.Elements))
	}
	for i, el := range result.Elements {
		testLiteralObject(t, el, expected[i])
	}
}

type ExpectedFunction struct {
	String string
	Type   object.FunctionObjectType
}

func testFunctionObject(t *testing.T, obj object.Object, expected ExpectedFunction) {
	if !object.IsFunction(obj) {
		t.Errorf("object is not Function. got=%T (%+v)", obj, obj)
	}

	result, isFunction := object.UnwrapReferenceObject(obj).(*object.Function)

	if !isFunction {
		t.Errorf("object is not Function. got=%T (%+v)", obj, obj)
	}

	if !testFunctionType(t, obj.Type(), expected.Type) {
		t.Errorf("function type incorrect. got=%s, want=%s",
			obj.Type().Signature(), expected.Type.Signature())
	}

	if result.Inspect() != expected.String {
		t.Errorf("function body incorrect. got=\n%s\n, want=\n%s",
			result.Inspect(), expected.String)
	}
}

func testFunctionType(t *testing.T, objectType object.ObjectType, expected object.FunctionObjectType) bool {
	functionType, ok := object.UnwrapReferenceType(objectType).(*object.FunctionObjectType)
	if !ok {
		t.Errorf("objectType is not function. got=%s", objectType.Signature())
		return false
	}

	for index, pt := range functionType.ParameterTypes {
		expectedSignature := expected.ParameterTypes[index].Signature()
		if pt.Signature() != expectedSignature {
			t.Errorf("function parameter %d has wrong type. got=%s, want=%s", index, objectType.Signature(), expectedSignature)
			return false
		}
	}

	if functionType.ReturnValueType.Signature() != expected.ReturnValueType.Signature() {
		t.Errorf("function return value has wrong type. got=%s, want=%s",
			functionType.ReturnValueType.Signature(), expected.ReturnValueType.Signature())
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	if !object.IsBoolean(obj) {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
	}

	result, isBoolean := object.UnwrapReferenceObject(obj).(*object.Boolean)

	if !isBoolean {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
	}
}

func testLiteralObject(t *testing.T, obj object.Object, expected interface{}) {
	switch expected := expected.(type) {
	case *big.Int:
		testIntegerObject(t, obj, expected)
	case string:
		testStringObject(t, obj, expected)
	case bool:
		testBooleanObject(t, obj, expected)
	case []interface{}:
		testArrayObject(t, obj, expected)
	case ExpectedFunction:
		testFunctionObject(t, obj, expected)
	default:
		t.Fatalf("unsupported expected type %T in test", expected)
	}
}
