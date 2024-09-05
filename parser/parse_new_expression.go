package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseNewExpression() ast.Expression {
	newExpr := &ast.NewExpression{Token: p.curToken}

	if !p.curTokenIs(token.New) {
		p.newError(nil, `Expected "new" keyword. got=%s`, p.curToken.Literal)
		return nil
	}

	p.nextToken()
	newExpr.Type = p.parseType()

	if newExpr.Type == nil {
		return nil
	}

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	newExpr.InitializationList = p.parseExpressionList(token.RightBrace)

	if newExpr.InitializationList == nil {
		return nil
	}

	return newExpr
}
