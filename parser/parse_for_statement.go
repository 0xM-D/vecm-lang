package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseForStatement() ast.Statement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()

	// Initialization
	if !p.curTokenIs(token.SEMICOLON) {
		stmt.Initialization = p.parseStatement()
	}
	if p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	} else {
		p.newError(stmt, "Expected ;")
		return nil
	}

	// Condition
	if !p.curTokenIs(token.SEMICOLON) {
		stmt.Condition = p.parseStatement()
	}

	// Afterthought
	if !p.peekTokenIs(token.RPAREN) {
		if p.curTokenIs(token.SEMICOLON) {
			p.nextToken()
		} else {
			p.newError(stmt, "Expected ;")
			return nil
		}
		stmt.AfterThought = p.parseStatement()
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	p.nextToken()
	if p.curTokenIs(token.LBRACE) {
		stmt.Body = p.parseBlockStatement()
	} else {
		stmt.Body = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
	}

	return stmt
}
