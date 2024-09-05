package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseIfStatement() ast.Statement {
	stmt := &ast.IfStatement{Token: p.curToken}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	p.nextToken()
	if p.curTokenIs(token.LeftBrace) {
		stmt.Consequence = p.parseBlockStatement()
	} else {
		stmt.Consequence = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
	}

	if p.peekTokenIs(token.Else) {
		p.nextToken()
		p.nextToken()
		if p.curTokenIs(token.LeftBrace) {
			stmt.Alternative = p.parseBlockStatement()
		} else {
			stmt.Alternative = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
		}
	}

	return stmt
}
