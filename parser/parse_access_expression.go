package parser

import (
	"github.com/DustTheory/interpreter/ast"
)

func (p *Parser) parseAccessExpression(left ast.Expression) ast.Expression {
	exp := &ast.AccessExpression{Token: p.curToken, Left: left, Right: nil}

	p.nextToken()
	exp.Right = p.parseIdentifier()

	return exp
}
