package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseNewExpression() ast.Expression {
	newExpr := &ast.NewExpression{Token: p.curToken}

	if !p.curTokenIs(token.NEW) {
		p.newError(nil, `Expected "new" keyword. got=%s`, p.curToken.Literal)
		return nil
	}

	p.nextToken()
	newExpr.Type = p.parseType()

	if newExpr.Type == nil {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	newExpr.InitializationList = p.parseExpressionList(token.RBRACE)

	if newExpr.InitializationList == nil {
		return nil
	}

	return newExpr
}
