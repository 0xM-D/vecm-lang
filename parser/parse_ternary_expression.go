package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseTernaryExpression(left ast.Expression) ast.Expression {
	ternaryToken := p.curToken
	p.nextToken() // Swallow ? token

	valueIfTrue := p.parseExpression(TernaryIf)

	if valueIfTrue == nil {
		return nil
	}

	if !p.expectPeek(token.Colon) {
		return nil
	}
	p.nextToken()

	valueIfFalse := p.parseExpression(Lowest)

	if valueIfFalse == nil {
		return nil
	}

	return &ast.TernaryExpression{
		Token:        ternaryToken,
		Condition:    left,
		ValueIfTrue:  valueIfTrue,
		ValueIfFalse: valueIfFalse,
	}
}
