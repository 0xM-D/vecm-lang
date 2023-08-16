package parser

import (
	"github.com/0xM-D/interpreter/ast"
)

func (p *Parser) parseAccessExpression(left ast.Expression) ast.Expression {
	exp := &ast.AccessExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Right = p.parseIdentifier()

	return exp
}
