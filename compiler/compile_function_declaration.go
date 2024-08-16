package compiler

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/context"
	"github.com/0xM-D/interpreter/util"
	"github.com/llir/llvm/ir"
)

func (c *Compiler) compileFunctionDeclaration(stmt *ast.FunctionDeclarationStatement, ctx *context.GlobalContext) {
	retType, err := util.GetLLVMType(stmt.Type.ReturnType)
	if err != nil {
		c.newCompilerError(stmt, err.Error())
		return;
	}

	paramTypes, err := getFunctionParamTypes(stmt)
	if err != nil {
		c.newCompilerError(stmt, err.Error())
		return;
	}

	name := stmt.Name.Value

	fn := context.NewFunctionContext(ctx, ctx.DeclareFunction(name, retType, paramTypes...), stmt.Parameters, stmt.Type.ParameterTypes)
	c.compileFunctionBody(stmt, fn)
}

func (c *Compiler) compileFunctionBody(stmt *ast.FunctionDeclarationStatement, fn *context.FunctionContext) {
	entryBlock := context.NewBlockContext(fn, fn.NewBlock(""))
	c.compileBlockStatement(stmt.Body, entryBlock)
}

func getFunctionParamTypes(stmt *ast.FunctionDeclarationStatement) ([]*ir.Param, error) {
	params := make([]*ir.Param, len(stmt.Type.ParameterTypes))
	for i, param := range stmt.Type.ParameterTypes {
		paramName := stmt.Parameters[i].Value
		paramType, err := util.GetLLVMType(param)
		if err != nil {
			return nil, err
		}
		params[i] = &ir.Param{LocalIdent: ir.NewLocalIdent(paramName), Typ: paramType}
	}
	return params, nil
}
