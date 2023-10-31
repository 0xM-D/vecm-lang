package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	p.nextToken()
	if p.curTokenIs(token.LBRACE) {
		expression.Consequence = p.parseBlockStatement()
	} else {
		expression.Consequence = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
	}

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		p.nextToken()
		if p.curTokenIs(token.LBRACE) {
			expression.Alternative = p.parseBlockStatement()
		} else {
			expression.Alternative = &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
		}
	}

	return expression
}
