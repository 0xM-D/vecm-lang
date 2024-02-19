package compiler

import (
	"github.com/0xM-D/interpreter/module"
)

type Compiler struct {
	Modules     map[string]*module.Module
	EntryModule *module.Module
	Errors []CompilerError
}

func InitializeCompiler() (*Compiler, error) {
	return &Compiler{Modules: map[string]*module.Module{}}, nil
}

func (c *Compiler) CompileModule(m *module.Module) string {
	ctx := c.compileProgram(m.Program)

	if(c.hasCompilerErrors()) {
		c.printCompilerErrors()
	}
	
	return ctx.Module.String()
}