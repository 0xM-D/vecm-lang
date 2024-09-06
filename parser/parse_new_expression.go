package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseNewExpression() ast.Expression {
	newToken := p.curToken
	if !p.curTokenIs(token.New) {
		p.newErrorf(nil, `Expected "new" keyword. got=%s`, p.curToken.Literal)
		return nil
	}

	p.nextToken()
	newType := p.parseType()

	if newType == nil {
		return nil
	}

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	initializationList := p.parseExpressionList(token.RightBrace)
	if initializationList == nil {
		return nil
	}

	return &ast.NewExpression{
		Token:              newToken,
		Type:               newType,
		InitializationList: initializationList,
	}
}
