package module

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/object"
	"github.com/DustTheory/interpreter/parser"
)

type Module struct {
	ModuleKey       string
	RootEnvironment object.Environment
	Program         *ast.Program
}

func ParseModule(moduleKey, code string) (*Module, []string) {
	l := lexer.New(string(code))
	p := parser.New(l)

	program := p.ParseProgram()

	module := &Module{
		ModuleKey:       moduleKey,
		RootEnvironment: *object.NewEnvironment(),
		Program:         program,
	}

	return module, p.Errors()
}
