package ast_test

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				ast.DeclarationStatement{
					Token: token.Token{Type: token.LET, Literal: "let", Linen: 0, Coln: 0},
					Name: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "myVar", Linen: 0, Coln: 0},
						Value: "myVar",
					},
					Value: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "anotherVar", Linen: 0, Coln: 0},
						Value: "anotherVar",
					},
					IsConstant: false,
					Type:       nil,
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String wrong. got=%q", program.String())
	}
}
