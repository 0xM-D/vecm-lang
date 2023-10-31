package parser

import "github.com/0xM-D/interpreter/ast"

func (p *Parser) prasePairExpression(left ast.Expression) ast.Expression {
	expr := &ast.PairExpression{Token: p.curToken, Left: left}

	p.nextToken() // :

	expr.Right = p.parseExpression(LOWEST)

	if expr.Right == nil {
		return nil
	}

	return expr
}
