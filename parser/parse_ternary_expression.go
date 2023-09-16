package parser

import (
	"github.com/0xM-D/interpreter/ast"
)

func (p *Parser) parseTernaryExpression(left ast.Expression) ast.Expression {
	expr := &ast.TernaryExpression{Token: p.curToken, Condition: left}

	p.nextToken() // Swallow ? token

	expr.TernaryValueExpression = p.parseExpression(LOWEST)
	if expr.TernaryValueExpression == nil {
		return nil
	}

	return expr
}
