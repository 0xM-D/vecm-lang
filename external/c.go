package external

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
	"github.com/DustTheory/interpreter/util"
	"github.com/pkg/errors"
)

func RunClangToLLVMIR(headerContent string) (string, error) {
	cmd := exec.Command("clang", "-x", "c", "-emit-llvm", "-S", "-fsyntax-only", "-")

	cmd.Stdin = bytes.NewReader([]byte(headerContent))

	var output bytes.Buffer
	cmd.Stdout = &output

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running clang for LLVM IR: %v\n%s", err, stderr.String())
	}

	return output.String(), nil
}

func GetCLibraryFunctionsFromPath(paths []string, token token.Token) ([]ast.FunctionDeclarationStatement, error) {
	// Create the `#include` directives for each header file
	clangInput, err := createCPathImportDirective(paths)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create C import directive")
	}

	// Compile the `#include` directives to LLVM IR
	clangIrOutput, err := compileCImportDirectives(clangInput)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile C import directives")
	}

	println(clangIrOutput)

	// Parse llvm IR to get function signatures
	funcs, err := ParseASTForFunctionDeclarations(clangIrOutput, token)

	if err != nil {
		return nil, errors.Wrap(err, "failed to import LLVM library")
	}

	println(funcs)

	return nil, nil
}

func ImportCStdlib(paths []string, token token.Token) ([]ast.FunctionDeclarationStatement, error) {
	// Create the `#include` directives for each header file
	clangInput := createCStdlibImportDirective(paths)

	// Compile the `#include` directives to LLVM IR
	clangIrOutput, err := compileCImportDirectives(clangInput)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile C import directives")
	}

	println(clangIrOutput)

	// Parse llvm IR to get function signatures
	library, err := ParseLLVMForFunctionDeclarations(clangIrOutput, token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to import LLVM library")
	}

	return library, nil
}

// ([]ast.FunctionDeclarationStatement, error)
func ParseASTForFunctionDeclarations(astDump string, token token.Token) ([]string, error) {
	var functions []string
	funcRegex := regexp.MustCompile(`FunctionDecl .*? ([a-zA-Z_][a-zA-Z0-9_]*) '([^(]+)\(([^)]*)\)'`)

	lines := strings.Split(astDump, "\n")
	for _, line := range lines {
		match := funcRegex.FindStringSubmatch(line)
		if match != nil {
			funcName := match[1]
			returnType := match[2]
			params := match[3]

			println(funcName, returnType, params)

			// functionParamNames := []*ast.Identifier{}
			// functionParamTypes := []ast.Type{}

			// // Convert C types to LLVM IR types
			llvmReturnType, err := util.GetVecmTypeFromCType(returnType, token)

			if err != nil {
				return nil, errors.Wrap(err, "failed to get LLVM type")
			}

			llvmParams := []string{} // convertParamsToLLVM(params)

			// Build the LLVM function declaration
			llvmFuncDecl := fmt.Sprintf("declare %s @%s(%s)", llvmReturnType, funcName, llvmParams)
			functions = append(functions, llvmFuncDecl)
		}
	}

	return functions, nil
}

func createCPathImportDirective(paths []string) (bytes.Buffer, error) {
	// Create the `#include` directives for each header file
	var input bytes.Buffer
	for _, path := range paths {
		absPath, err := filepath.Abs(path) // Convert to absolute path for clarity
		if err != nil {
			return input, errors.Wrapf(err, "failed to resolve absolute path for %s", path)
		}
		input.WriteString(fmt.Sprintf("#include \"%s\"\n", absPath))
	}

	return input, nil
}

func createCStdlibImportDirective(headers []string) bytes.Buffer {
	// Create the `#include` directives for each header file
	var input bytes.Buffer
	for _, header := range headers {
		input.WriteString(fmt.Sprintf("#include <%s>\n", header))
	}

	return input
}

// func compileCImportDirectives(input bytes.Buffer) (string, error) {
// 	cmd := exec.Command("clang", "-x", "c", "-fkeep-static-consts", "-Xclang", "-ast-dump", "-fsyntax-only", "-")

// 	cmd.Stdin = &input

// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	err := cmd.Run()
// 	if err != nil {
// 		return "", errors.Wrap(err, "failed to compile C import directives")
// 	}

// 	return out.String(), nil
// }

func compileCImportDirectives(input bytes.Buffer) (string, error) {
	cmd := exec.Command("clang", "-x", "c", "-emit-llvm", "-S", "-")

	cmd.Stdin = &input

	var output bytes.Buffer
	cmd.Stdout = &output

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running clang for LLVM IR: %v\n%s", err, stderr.String())
	}

	return output.String(), nil
}
