package runtime

import (
	"errors"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalExportStatement(node *ast.ExportStatement, env *object.Environment) (object.Object, error) {
	stmtResult, err := r.Eval(node.Statement, env)
	if err != nil {
		return nil, err
	}

	variableReference, ok := stmtResult.(*object.VariableReference)
	if !ok {
		return nil, errors.New("export statement expects a variable reference")
	}

	env.GetStore()[variableReference.Name].IsExported = true

	return variableReference, nil
}
