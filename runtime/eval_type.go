package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalType(typeNode ast.Type, env *object.Environment) (object.ObjectType, error) {
	switch casted := typeNode.(type) {
	case ast.HashType:
		keyType, err := r.evalType(casted.KeyType, env)
		if err != nil {
			return nil, err
		}
		valueType, err := r.evalType(casted.ValueType, env)
		if err != nil {
			return nil, err
		}
		return &object.HashObjectType{KeyType: keyType, ValueType: valueType}, nil
	case ast.ArrayType:
		elementType, err := r.evalType(casted.ElementType, env)
		if err != nil {
			return nil, err
		}
		return &object.ArrayObjectType{ElementType: elementType}, nil
	case ast.NamedType:
		namedType, found := env.GetObjectType(casted.TokenLiteral())
		if !found {
			return nil, fmt.Errorf("unknown type %s in: %s", casted.TokenLiteral(), typeNode.String())
		}
		return namedType, nil
	case ast.FunctionType:
		parameterTypes := []object.ObjectType{}
		returnType, returnTypeErr := r.evalType(casted.ReturnType, env)
		if returnTypeErr != nil {
			return nil, returnTypeErr
		}

		for _, p := range casted.ParameterTypes {
			parsedType, parsedTypeError := r.evalType(p, env)
			if parsedTypeError != nil {
				return nil, parsedTypeError
			}
			parameterTypes = append(parameterTypes, parsedType)
		}
		return &object.FunctionObjectType{ParameterTypes: parameterTypes, ReturnValueType: returnType}, nil
	}

	return nil, fmt.Errorf("unknown type: %s", typeNode.String())
}
