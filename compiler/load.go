package compiler

import (
	"github.com/0xM-D/interpreter/module"
	"github.com/0xM-D/interpreter/util"
)

func (c *Compiler) LoadEntryModuleFromFile(filepath string) (*module.Module, error) {
	module, err := c.LoadModuleFromFile(filepath)
	if err != nil {
		return nil, err
	}

	c.EntryModule = module
	return module, nil
}

func (c *Compiler) LoadModuleFromFile(filepath string) (*module.Module, error) {
	file, err := util.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return c.LoadModule(file.Path, string(file.Bytes))
}

func (c *Compiler) LoadModule(moduleKey, code string) (*module.Module, error) {
	module := module.ParseModule(moduleKey, code)
	parserErrors := module.Parser.Errors()
	if len(parserErrors) != 0 {
		return nil, getConcatenatedParserErrors(parserErrors)
	}

	c.Modules[moduleKey] = module

	ir := c.CompileModule(module)

	println(ir)

	return module, nil
}
