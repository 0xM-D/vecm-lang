package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
	"github.com/DustTheory/interpreter/object"
	"github.com/DustTheory/interpreter/util"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

func (c *Compiler) compileFunctionDeclaration(stmt *ast.FunctionDeclarationStatement, ctx *context.GlobalContext) {
	retType, err := util.GetLLVMType(stmt.Type.ReturnType)
	if err != nil {
		c.newCompilerError(stmt, "%e", err)
		return
	}

	paramTypes, err := getFunctionParamTypes(stmt)
	if err != nil {
		c.newCompilerError(stmt, "%e", err)
		return
	}

	name := stmt.Name.Value

	function, err := ctx.DeclareFunction(
		name,
		retType,
		paramTypes...,
	)

	if err != nil {
		c.newCompilerError(stmt, "%e", err)
		return
	}

	fnCtx := context.NewFunctionContext(
		ctx,
		function,
		stmt.Parameters,
		stmt.Type.ParameterTypes,
	)

	c.compileFunctionBody(stmt, fnCtx)
}

func (c *Compiler) compileFunctionBody(stmt *ast.FunctionDeclarationStatement, fn *context.FunctionContext) {
	bodyBlock := context.NewBlockContext(fn, fn.NewBlock(""))
	c.compileBlockStatement(stmt.Body, bodyBlock)

	if bodyBlock.Block.Term == nil {
		if stmt.Type.ReturnType.String() == object.VoidKind.Signature() {
			bodyBlock.NewRet(nil)
		} else {
			c.newCompilerError(stmt, "Function must return a value")
		}
	}
}

func getFunctionParamTypes(stmt *ast.FunctionDeclarationStatement) ([]*ir.Param, error) {
	params := make([]*ir.Param, len(stmt.Type.ParameterTypes))
	for i, param := range stmt.Type.ParameterTypes {
		paramName := stmt.Parameters[i].Value
		paramType, err := util.GetLLVMType(param)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to get LLVM type")
		}
		params[i] = &ir.Param{LocalIdent: ir.NewLocalIdent(paramName), Typ: paramType}
	}
	return params, nil
}
