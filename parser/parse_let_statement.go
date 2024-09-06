package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	declStmtToken := p.curToken
	if !p.expectPeek(token.Ident) {
		return nil
	}

	nameIdentifier := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.Assign) {
		return nil
	}

	p.nextToken()

	value := p.parseExpression(Lowest)
	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return &ast.LetStatement{
		DeclarationStatement: ast.DeclarationStatement{
			Token:      declStmtToken,
			Name:       nameIdentifier,
			Value:      value,
			IsConstant: false,
			Type:       nil,
		},
	}
}
