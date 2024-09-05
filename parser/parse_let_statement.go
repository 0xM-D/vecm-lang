package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.DeclarationStatement{Token: p.curToken}
	if !p.expectPeek(token.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.Assign) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return &ast.LetStatement{DeclarationStatement: *stmt}
}
