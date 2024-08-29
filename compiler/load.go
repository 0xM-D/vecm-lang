package compiler

import (
	"fmt"

	"github.com/DustTheory/interpreter/module"
)

func (c *Compiler) loadModule(moduleKey, code string) (*module.Module, bool) {
	module, failedToLoad := module.ParseModule(moduleKey, code)

	if len(failedToLoad) != 0 {
		printParserErrors(failedToLoad)
	}

	c.Modules[moduleKey] = module
	return module, false
}

func printParserErrors(errors []string) {
	fmt.Printf("parser has %d errors\n", len(errors))
	for _, msg := range errors {
		fmt.Printf("parser error: %s", msg)
	}
}
