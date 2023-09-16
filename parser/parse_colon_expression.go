package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseColonExpression(left ast.Expression) ast.Expression {
	expr := &ast.ColonExpression{Token: p.curToken, Left: left}

	if !p.curTokenIs(token.COLON) {
		p.newError(nil, "Expected :")
		return nil
	}

	p.nextToken() // Swallow : token

	expr.Right = p.parseExpression(LOWEST)
	if expr.Right == nil {
		return nil
	}

	return expr

}
