package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalImportStatement(node *ast.ImportStatement, env *object.Environment) (object.Object, error) {
	module, failedToLoad := r.loadModuleFromFile(node.ImportPath)

	if failedToLoad {
		return nil, fmt.Errorf("failed to load module: %s", node.ImportPath)
	}

	store := module.RootEnvironment.GetStore()

	for _, identifier := range node.ImportedIdentifiers {
		storeEntry, found := store[identifier.Value]
		if !found || !storeEntry.IsExported {
			return nil, fmt.Errorf("imported name %s not found in exports", identifier.Value)
		}

		_, err := env.Declare(identifier.Value, storeEntry.IsConstant, storeEntry.Object)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
