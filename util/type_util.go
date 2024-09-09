package util

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
	"github.com/llir/llvm/ir/types"
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
	// TODO: As we add more types, we need to add more cases here
	//
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
