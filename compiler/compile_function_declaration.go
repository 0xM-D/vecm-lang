package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) compileFunctionDeclaration(stmt *ast.FunctionDeclarationStatement, m *ir.Module) error {
	retType, err := getLLVMType(stmt.Type.ReturnType)
	if err != nil {
		return err
	}
	paramTypes, err := getFunctionParamTypes(stmt)
	if err != nil {
		return err
	}
	fn := m.NewFunc(stmt.Name.Value, retType, paramTypes...)
	compileFunctionBody(stmt, fn)
	return nil
}

func getFunctionParamTypes(stmt *ast.FunctionDeclarationStatement) ([]*ir.Param, error) {
	params := make([]*ir.Param, len(stmt.Type.ParameterTypes))
	for i, param := range stmt.Type.ParameterTypes {
		paramName := stmt.Parameters[i].Value
		paramType, err := getLLVMType(param)
		if err != nil {
			return nil, err
		}
		params[i] = &ir.Param{LocalIdent: ir.NewLocalIdent(paramName), Typ: paramType}
	}
	return params, nil
}

func compileFunctionBody(stmt *ast.FunctionDeclarationStatement, fn *ir.Func) {
	entryBlock := fn.NewBlock("")
	for _, stmt := range stmt.Body.Statements {
		compileStatement(stmt, entryBlock)
	}
}

func compileStatement(stmt ast.Statement, b *ir.Block) {
	switch stmt := stmt.(type) {
	case *ast.ReturnStatement:
		compileReturnStatement(stmt, b)
	}
}

func compileReturnStatement(stmt *ast.ReturnStatement, b *ir.Block) {
	value := compileExpression(stmt.ReturnValue)
	b.NewRet(value)
}

func compileExpression(expr ast.Expression) value.Value {
	return nil
}
