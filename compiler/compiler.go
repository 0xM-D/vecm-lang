package compiler

import (
	"bytes"
	"fmt"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/module"
	"github.com/0xM-D/interpreter/util"
)

type Compiler struct {
	Modules     map[string]*module.Module
	EntryModule *module.Module
}

func InitializeCompiler() (*Compiler, error) {
	return &Compiler{Modules: map[string]*module.Module{}}, nil
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

	ir, err := c.Compile(module)

	if err != nil {
		println(err.Error())
	}
	println(ir, err)

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

func (c *Compiler) Compile(m *module.Module) (string, error) {
	// ir_module := ir.NewModule()
	// fn := ir_module.NewFunc("main", types.I32, ir.NewParam("argc", types.I32), ir.NewParam("argv", types.NewPointer(types.I8Ptr)))
	// entry := fn.NewBlock("entry")
	// blk := fn.NewBlock("")
	// blk.NewRet(constant.NewInt(types.I32, 0))
	// entry.NewBr(blk)
	// return ir_module.String(), nil

	govno, err := c.Compi(m.Program)

	if err != nil {
		return "", err
	}

	return govno.String(), err
}

func (c *Compiler) Compi(node ast.Node) (Govno, error) {
	switch node := node.(type) {
	case *ast.Program:
		return c.compileProgram(node)
	default:
		return nil, newError("Invalid top level statement: %T", node)
	}

	return nil, nil
}
