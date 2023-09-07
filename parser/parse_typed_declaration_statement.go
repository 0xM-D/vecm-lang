package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseTypedDeclarationStatement() *ast.TypedDeclarationStatement {
	stmt := &ast.DeclarationStatement{Token: p.curToken}

	stmt.Type = p.parseType()
	p.nextToken()

	stmt.Name = p.parseIdentifier().(*ast.Identifier)

	if p.peekToken.Type != token.ASSIGN {
		p.newError(stmt, "invalid token in typed declaration statement. expected=%q got=%q", "=", p.peekToken.Literal)
		return nil
	}

	p.nextToken()
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &ast.TypedDeclarationStatement{DeclarationStatement: *stmt}
}
