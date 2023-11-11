package runtime

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/object"
	"github.com/0xM-D/interpreter/parser"
)

type Module struct {
	ModuleKey       string
	RootEnvironment object.Environment
	Lexer           *lexer.Lexer
	Parser          *parser.Parser
	Program         *ast.Program
}

func ImportModule(moduleKey, code string) *Module {
	l := lexer.New(string(code))
	p := parser.New(l)

	program := p.ParseProgram()

	module := &Module{
		ModuleKey:       moduleKey,
		RootEnvironment: *object.NewEnvironment(),
		Lexer:           l,
		Parser:          p,
		Program:         program,
	}

	return module
}
