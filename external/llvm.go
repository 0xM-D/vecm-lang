package external

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
	"github.com/DustTheory/interpreter/util"
	"github.com/llir/llvm/asm"
	"github.com/pkg/errors"
)

func ParseLLVMForFunctionDeclarations(llvmIr string, token token.Token) ([]ast.FunctionDeclarationStatement, error) {
	module, parseErr := asm.ParseString("", llvmIr)
	if parseErr != nil {
		return nil, errors.Wrap(parseErr, "failed to parse LLVM IR")
	}

	funcs := []ast.FunctionDeclarationStatement{}
	for _, f := range module.Funcs {
		functionParamNames := []*ast.Identifier{}
		functionParamTypes := []ast.Type{}
		for _, arg := range f.Params {
			functionParamNames = append(functionParamNames, &ast.Identifier{
				Token: token,
				Value: arg.LocalName,
			})

			paramType, err := util.GetVecmTypeFromLLVMType(arg.Type(), token)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get LLVM type")
			}

			functionParamTypes = append(functionParamTypes, paramType)
		}

		returnType, err := util.GetVecmTypeFromLLVMType(f.Sig.RetType, token)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get LLVM type")
		}

		funcs = append(funcs, ast.FunctionDeclarationStatement{
			Token:      token,
			Name:       &ast.Identifier{Token: token, Value: f.Name()},
			Parameters: functionParamNames,
			Type:       ast.FunctionType{Token: token, ReturnType: returnType, ParameterTypes: functionParamTypes},
			Body:       nil,
		})
	}

	return funcs, nil
}
