package parser

import (
	"github.com/DustTheory/interpreter/ast"
)

func (p *Parser) parseCLangStatement() *ast.CLangStatement {
	clangToken := p.curToken

	clangCode := p.l.ExternCode()
	p.nextToken() // Swallow the CLang code

	return &ast.CLangStatement{
		Token:     clangToken,
		CLangCode: clangCode,
	}
}
