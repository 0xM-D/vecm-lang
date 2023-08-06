package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseType() ast.Type {

	var result ast.Type

	switch p.curToken.Type {
	case token.MAP_TYPE:
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		p.nextToken()

		keyType := p.parseType()

		if !p.expectPeek(token.DASH_ARROW) {
			return nil
		}
		p.nextToken()

		valueType := p.parseType()

		if !p.expectPeek(token.RBRACE) {
			return nil
		}

		result = ast.HashType{Token: p.curToken, KeyType: keyType, ValueType: valueType}
	case token.IDENT:
		typeIdentifier := p.parseIdentifier().(*ast.Identifier)
		result = ast.NamedType{Token: p.curToken, TypeName: *typeIdentifier}
	}

	for p.peekToken.Type == token.ARRAY_TYPE {
		p.nextToken()
		result = ast.ArrayType{Token: p.curToken, ElementType: result}
	}

	return result
}
