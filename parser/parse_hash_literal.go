package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()

		expr := p.parseExpression(LOWEST)

		if expr == nil {
			return nil
		}

		colonExpression, ok := expr.(*ast.ColonExpression)

		if !ok {
			p.newError(colonExpression, "Expected colon expression. got=%T", expr)
			return nil
		}

		hash.Pairs[colonExpression.Left] = colonExpression.Right
		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}

	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return hash
}
