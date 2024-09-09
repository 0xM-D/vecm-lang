package compiler

import (
	"github.com/DustTheory/interpreter/module"
)

type Compiler struct {
	Modules map[string]*module.Module

	EntryModule *module.Module
	Errors      []Error
}

func New() (*Compiler, error) {
	return &Compiler{
		Modules:     map[string]*module.Module{},
		EntryModule: nil,
		Errors:      []Error{},
	}, nil
}

func (c *Compiler) LoadModule(moduleKey, code string) (*module.Module, bool) {
	module, failedToLoad := c.loadModule(moduleKey, code)
	return module, failedToLoad
}

func (c *Compiler) CompileModule(moduleKey string) (string, bool) {
	module := c.Modules[moduleKey]
	ctx := c.compileProgram(module.Program)

	return ctx.Module.String(), c.hasCompilerErrors()
}
