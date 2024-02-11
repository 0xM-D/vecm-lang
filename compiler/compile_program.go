package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileProgram(program *ast.Program) (Govno, error) {
	m := ir.NewModule()

	for _, statement := range program.Statements {
		err := c.compileTopLevelStatement(statement, m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (c *Compiler) compileTopLevelStatement(statement ast.Statement, m *ir.Module) error {
	switch statement := statement.(type) {
	case *ast.FunctionDeclarationStatement:
		return c.compileFunctionDeclaration(statement, m)
	case *ast.ExportStatement:
		return c.compileExportStatement(statement, m)
	// case *ast.TypedDeclarationStatement:
	// 	return nil // TODO
	// case *ast.DeclarationStatement:
	// 	return nil // TODO
	default:
		return newError("Invalid statement on top level: %T", statement)
	}
}
