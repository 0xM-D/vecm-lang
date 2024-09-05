package runtime

import (
	"fmt"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalImportStatement(node *ast.ImportStatement, env *object.Environment) error {
	module, failedToLoad := r.loadModuleFromFile(node.ImportPath)

	if failedToLoad {
		return fmt.Errorf("failed to load module: %s", node.ImportPath)
	}

	store := module.RootEnvironment.GetStore()

	// print store
	println(len(store))
	for key, value := range store {
		fmt.Printf("key: %s, value: %v\n", key, value)
	}

	for _, identifier := range node.ImportedIdentifiers {
		storeEntry, found := store[identifier.Value]
		if !found || !storeEntry.IsExported {
			return fmt.Errorf("imported name %s not found in exports", identifier.Value)
		}

		_, err := env.Declare(identifier.Value, storeEntry.IsConstant, storeEntry.Object)
		if err != nil {
			return fmt.Errorf("error in import statement %s: %w", identifier.Value, err)
		}
	}

	return nil
}
