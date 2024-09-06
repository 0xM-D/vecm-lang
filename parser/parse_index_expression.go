package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	// Swallow the left bracket
	indexExpressionToken := p.curToken
	p.nextToken()

	index := p.parseExpression(Lowest)

	if !p.expectPeek(token.RightBracket) {
		return nil
	}

	return &ast.IndexExpression{
		Token: indexExpressionToken,
		Left:  left,
		Index: index,
	}
}
