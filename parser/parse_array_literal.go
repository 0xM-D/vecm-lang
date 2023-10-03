package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}

	array.Type = p.parseType()
	if array.Type == nil {
		return nil
	}

	return p.parseArrayLiteralRest(array)
}

func (p *Parser) parseArrayLiteralRest(array *ast.ArrayLiteral) ast.Expression {
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	array.Elements = p.parseExpressionList(token.RBRACE)
	return array
}
