package compiler

import (
	"bytes"
	"fmt"

	"github.com/0xM-D/interpreter/module"
	"github.com/0xM-D/interpreter/util"
)

type Compiler struct {
	Modules     map[string]*module.Module
	EntryModule *module.Module
}

func InitializeCompiler() (*Compiler, error) {
	return &Compiler{}, nil
}

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

	// c.Compile(module)

	return module, nil
}

func getConcatenatedParserErrors(errors []string) error {
	var errs bytes.Buffer

	errs.WriteString(fmt.Sprintf("parser has %d errors\n", len(errors)))
	for _, msg := range errors {
		errs.WriteString(fmt.Sprintf("parser error: %s", msg))
	}

	return fmt.Errorf(errs.String())
}
