package evaluator_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/evaluator"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/object"
	"github.com/0xM-D/interpreter/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != evaluator.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	if !object.IsInteger(obj) {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	result := object.UnwrapReferenceObject(obj).(*object.Integer)

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	if !object.IsString(obj) {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}

	result := object.UnwrapReferenceObject(obj).(*object.String)

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q",
			result.Value, expected)
		return false
	}
	return true
}

func testArrayObject(t *testing.T, obj object.Object, expected []string) bool {
	if !object.IsArray(obj) {
		t.Errorf("object is not Array. got=%T (%+v)", obj, obj)
		return false
	}

	result := object.UnwrapReferenceObject(obj).(*object.Array)

	if len(result.Elements) != len(expected) {
		t.Errorf("Incorrect array length. expected=%d. got=%d", len(result.Elements), len(expected))
		return false
	}
	for i, el := range result.Elements {
		if el.Inspect() != expected[i] {
			t.Errorf("Array mismatch at index %d. expected=%s got=%s", i, el.Inspect(), expected[i])
			return false
		}
	}
	return true
}

type ExpectedFunction struct {
	String string
	Type   object.FunctionObjectType
}

func testFunctionObject(t *testing.T, obj object.Object, expected ExpectedFunction) bool {
	if !object.IsFunction(obj) {
		t.Errorf("object is not Function. got=%T (%+v)", obj, obj)
		return false
	}

	result := object.UnwrapReferenceObject(obj).(*object.Function)

	if !testFunctionType(t, obj.Type(), expected.Type) {

	}

	if result.Inspect() != expected.String {
		t.Errorf("function body incorrect. got=\n%s\n, want=\n%s",
			result.Inspect(), expected.String)
		return false
	}

	return true
}

func testFunctionType(t *testing.T, objectType object.ObjectType, expected object.FunctionObjectType) bool {
	functionType, ok := objectType.(*object.FunctionObjectType)
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
		t.Errorf("function return value has wrong type. got=%s, want=%s", functionType.ReturnValueType.Signature(), expected.ReturnValueType.Signature())
		return false
	}

	return true
}
