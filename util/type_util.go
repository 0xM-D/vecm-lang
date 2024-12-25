package util

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
	"github.com/DustTheory/interpreter/token"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

func GetLLVMType(t ast.Type) (types.Type, error) {
	switch t := t.(type) {
	case ast.NamedType:
		return GetIntrinsicLLVMType(t)
	case ast.VoidType:
		return &types.VoidType{}, nil
	}

	return nil, fmt.Errorf("type %s cannot be converted to appropriate LLVM type", t.String())
}

func GetIntrinsicLLVMType(t ast.NamedType) (types.Type, error) {
	// As we add more types, we need to add more cases here
	//nolint:exhaustive,mnd // We will handle all cases eventually, magic values are cleaner here
	switch object.Kind(t.TypeName.Value) {
	case "char":
		return types.NewInt(8), nil
	case "int":
		return types.NewInt(32), nil
	case object.Int8Kind:
		return types.NewInt(8), nil
	case object.Int16Kind:
		return types.NewInt(16), nil
	case object.Int32Kind:
		return types.NewInt(32), nil
	case object.Int64Kind:
		return types.NewInt(64), nil
	case object.UInt8Kind:
		return types.NewInt(8), nil
	case object.UInt16Kind:
		return types.NewInt(16), nil
	case object.UInt32Kind:
		return types.NewInt(32), nil
	case object.UInt64Kind:
		return types.NewInt(64), nil
	case object.BooleanKind:
		return types.NewInt(1), nil
	case object.Float32Kind:
		return &types.FloatType{Kind: types.FloatKindFloat}, nil
	case object.Float64Kind:
		return &types.FloatType{Kind: types.FloatKindDouble}, nil
	case object.VoidKind:
		return &types.VoidType{}, nil
	}
	return nil, fmt.Errorf("type %s (%T) cannot be converted to appropriate LLVM type", t.String(), t.TypeName)
}
func GetVecmTypeFromLLVMType(t types.Type, token token.Token) (ast.Type, error) {
	switch t := t.(type) {
	case *types.VoidType:
		return ast.VoidType{Token: token}, nil
	case *types.IntType:
		return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "int"}}, nil
	case *types.FloatType:
		switch t.Kind {
		case types.FloatKindFloat:
			return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "float"}}, nil
		case types.FloatKindDouble:
			return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "double"}}, nil
		case types.FloatKindFP128:
			return nil, errors.New("FP128 type not supported")
		case types.FloatKindX86_FP80:
			return nil, errors.New("X86_FP80 type not supported")
		case types.FloatKindPPC_FP128:
			return nil, errors.New("PPC_FP128 type not supported")
		case types.FloatKindHalf:
			return nil, errors.New("Half type not supported")
		}
	case *types.PointerType:
		return nil, errors.New("pointer type not supported")
	case *types.VectorType:
		return nil, errors.New("vector type not supported")
	case *types.ArrayType:
		elemType, err := GetVecmTypeFromLLVMType(t.ElemType, token)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get array element type")
		}
		return ast.ArrayType{Token: token, ElementType: elemType}, nil
	case *types.StructType:
		return nil, errors.New("struct type not supported")
	case *types.FuncType:
		return nil, errors.New("function type not supported")
	case *types.MMXType:
		return nil, errors.New("MMX type not supported")
	case *types.LabelType:
		return nil, errors.New("label type not supported")
	case *types.TokenType:
		return nil, errors.New("token type not supported")
	case *types.MetadataType:
		return nil, errors.New("metadata type not supported")
	}
	return nil, fmt.Errorf("type %s (%T) not supported", t.String(), t)
}

func GetVecmTypeFromCType(cType string, token token.Token) (ast.Type, error) {
	switch cType {
	case "int":
		return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "int"}}, nil
	case "void":
		return ast.VoidType{Token: token}, nil
	case "char":
		return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "i8"}}, nil
	case "float":
		return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "float"}}, nil
	case "double":
		return ast.NamedType{Token: token, TypeName: ast.Identifier{Token: token, Value: "double"}}, nil
	default:
		return nil, fmt.Errorf("type \"%s\" cannot be converted to appropriate LLVM type", cType)
	}
}
