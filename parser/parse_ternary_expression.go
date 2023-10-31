package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseTernaryExpression(left ast.Expression) ast.Expression {
	expr := &ast.TernaryExpression{Token: p.curToken, Condition: left}

	p.nextToken() // Swallow ? token

	expr.ValueIfTrue = p.parseExpression(TERNARY_IF)

	if expr.ValueIfTrue == nil {
		return nil
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}
	p.nextToken()

	expr.ValueIfFalse = p.parseExpression(LOWEST)

	if expr.ValueIfFalse == nil {
		return nil
	}

	return expr
}
