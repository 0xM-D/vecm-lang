package parser

import (
	"github.com/DustTheory/interpreter/ast"
)

func (p *Parser) parseIdentifier() *ast.Identifier {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIdentifierAsExpression() ast.Expression {
	return p.parseIdentifier()
}
