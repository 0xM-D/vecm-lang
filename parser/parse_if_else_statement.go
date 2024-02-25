package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseIfStatement() ast.Statement {
	stmt := &ast.IfStatement{Token: p.curToken}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	p.nextToken()
	if p.curTokenIs(token.LBRACE) {
		stmt.Consequence = p.parseBlockStatement()
	} else {
		stmt.Consequence = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
	}

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		p.nextToken()
		if p.curTokenIs(token.LBRACE) {
			stmt.Alternative = p.parseBlockStatement()
		} else {
			stmt.Alternative = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
		}
	}

	return stmt
}
