package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseDeclarationStatement(assignToken token.TokenType) *ast.DeclarationStatement {
	stmt := &ast.DeclarationStatement{Token: p.curToken}

	stmt.Name = p.parseIdentifier().(*ast.Identifier)

	if !p.peekTokenIs(assignToken) {
		p.newError(stmt, "invalid token in declaration statement. expected=%q got=%q", assignToken, p.peekToken.Literal)
		return nil
	}

	p.nextToken()
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
