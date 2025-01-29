package runtime

import (
	"errors"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalCLangStatement(node *ast.CLangStatement, env *object.Environment) (object.Object, error) {
	return nil, errors.New("not implemented")
}
