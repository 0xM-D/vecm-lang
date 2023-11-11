package evaluator

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) (object.Object, error) {
	var result object.Object
	var err error

	for _, statement := range block.Statements {
		result, err = Eval(statement, env)
		if err != nil {
			return nil, err
		}
		if result != nil && object.IsReturnValue(result) {
			return result, nil
		}
	}
	return result, nil
}
