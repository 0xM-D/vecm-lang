package parser

import "github.com/DustTheory/interpreter/ast"

func (p *Parser) prasePairExpression(left ast.Expression) ast.Expression {
	pairToken := p.curToken
	p.nextToken() // :

	right := p.parseExpression(Lowest)

	if right == nil {
		return nil
	}

	return &ast.PairExpression{
		Token: pairToken,
		Left:  left,
		Right: right,
	}
}
