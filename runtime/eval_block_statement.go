package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalBlockStatement(block *ast.BlockStatement, env *object.Environment) (object.Object, error) {
	var result object.Object
	var err error

	for _, statement := range block.Statements {
		result, err = r.Eval(statement, env)
		if err != nil {
			return nil, err
		}
		if result != nil && object.IsReturnValue(result) {
			return result, nil
		}
	}
	return result, nil
}
