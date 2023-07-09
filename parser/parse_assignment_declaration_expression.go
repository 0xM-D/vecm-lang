package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseAssignmentDeclarationStatement() *ast.AssignmentDeclarationStatement {
	stmt := &ast.AssignmentDeclarationStatement{Token: p.curToken}

	stmt.Name = p.parseIdentifier().(*ast.Identifier)
	p.nextToken()
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
