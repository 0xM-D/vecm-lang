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
		p.errors = append(p.errors, "Expected ;")
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
			p.errors = append(p.errors, "Expected ;")
			return nil
		}
		stmt.AfterThought = p.parseStatement()
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}
