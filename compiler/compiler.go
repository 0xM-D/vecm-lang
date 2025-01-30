package compiler

import (
	"github.com/DustTheory/interpreter/module"
	"github.com/llir/llvm/ir"
)

type Compiler struct {
	Modules map[string]*module.Module

	EntryModule *module.Module
	Errors      []Error
}

type IrModule struct {
	CoreModule      *ir.Module
	LinkedModulesIR []string
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

func (c *Compiler) CompileModule(moduleKey string) (IrModule, bool) {
	module := c.Modules[moduleKey]
	ctx := c.compileProgram(module.Program)

	return IrModule{
		CoreModule:      ctx.Module,
		LinkedModulesIR: ctx.LinkedModulesIR,
	}, c.hasCompilerErrors()
}

func (c *Compiler) AddModule(moduleKey string, module *module.Module) {
	c.Modules[moduleKey] = module
}
