package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseForStatement() ast.Statement {
	stmt := &ast.ForStatement{
		Token:          p.curToken,
		Initialization: nil,
		Condition:      nil,
		AfterThought:   nil,
		Body:           nil,
	}

	if !p.expectPeek(token.LeftParen) {
		return nil
	}
	p.nextToken()

	// Initialization
	if !p.curTokenIs(token.Semicolon) {
		stmt.Initialization = p.parseStatement()
	}
	if p.curTokenIs(token.Semicolon) {
		p.nextToken()
	} else {
		p.newErrorf(stmt, "Expected ;")
		return nil
	}

	// Condition
	if !p.curTokenIs(token.Semicolon) {
		stmt.Condition = p.parseStatement()
	}

	// Afterthought
	if !p.peekTokenIs(token.RightParen) {
		if p.curTokenIs(token.Semicolon) {
			p.nextToken()
		} else {
			p.newErrorf(stmt, "Expected ;")
			return nil
		}
		stmt.AfterThought = p.parseStatement()
	}

	if !p.expectPeek(token.RightParen) {
		return nil
	}

	p.nextToken()
	if p.curTokenIs(token.LeftBrace) {
		stmt.Body = p.parseBlockStatement()
	} else {
		stmt.Body = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
	}

	return stmt
}
