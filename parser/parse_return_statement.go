package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	if p.curTokenIs(token.Semicolon) {
		return stmt
	}

	stmt.ReturnValue = p.parseExpression(Lowest)
	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}
