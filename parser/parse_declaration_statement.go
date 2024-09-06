package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseDeclarationStatement(assignToken token.Type) *ast.DeclarationStatement {
	stmt := &ast.DeclarationStatement{Token: p.curToken, Name: nil, Value: nil, IsConstant: false, Type: nil}

	stmt.Name = p.parseIdentifier()

	if !p.peekTokenIs(assignToken) {
		p.newErrorf(stmt, "invalid token in declaration statement. expected=%q got=%q", assignToken, p.peekToken.Literal)
		return nil
	}

	p.nextToken()
	p.nextToken()

	stmt.Value = p.parseExpression(Lowest)

	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}
