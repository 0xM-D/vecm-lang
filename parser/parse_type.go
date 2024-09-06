package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseType() ast.Type {
	var result ast.Type

	if p.curToken.Type == token.ArrayType {
		p.nextToken()
		elementType := p.parseType()
		if elementType == nil {
			return nil
		}
		return ast.ArrayType{Token: p.curToken, ElementType: elementType}
	}

	//nolint:exhaustive // We don't need to handle all token types here
	switch p.curToken.Type {
	case token.MapType:
		result = p.parseMapType()
	case token.FunctionType:
		result = p.parseFunctionType()
	case token.Ident:
		typeIdentifier := p.parseIdentifier()
		result = ast.NamedType{Token: p.curToken, TypeName: *typeIdentifier}
	default:
		p.newErrorf(nil, `Expected type. got=%s`, p.curToken.Literal)
		return nil
	}
	return result
}

func (p *Parser) parseMapType() ast.Type {
	if !p.expectPeek(token.LeftBrace) {
		return nil
	}
	p.nextToken()

	keyType := p.parseType()

	if !p.expectPeek(token.DashArrow) {
		return nil
	}
	p.nextToken()

	valueType := p.parseType()

	if !p.expectPeek(token.RightBrace) {
		return nil
	}

	return ast.HashType{Token: p.curToken, KeyType: keyType, ValueType: valueType}
}

func (p *Parser) parseFunctionType() ast.Type {
	functionTypeToken := p.curToken

	parameterTypes := p.parseFunctionTypeParameters()

	if parameterTypes == nil {
		return nil
	}

	if !p.expectPeek(token.DashArrow) {
		return nil
	}
	p.nextToken()

	returnType := p.parseType()

	if returnType == nil {
		return nil
	}

	return ast.FunctionType{
		Token:          functionTypeToken,
		ParameterTypes: parameterTypes,
		ReturnType:     returnType,
	}
}

func (p *Parser) parseFunctionTypeParameters() []ast.Type {
	parameterTypes := []ast.Type{}

	if !p.expectPeek(token.LeftParen) {
		return nil
	}

	for {
		if p.peekTokenIs(token.RightParen) {
			break
		}
		p.nextToken()

		nextType := p.parseType()
		if nextType == nil {
			return nil
		}

		parameterTypes = append(parameterTypes, nextType)

		if !p.peekTokenIs(token.Comma) {
			break
		}
		p.nextToken()
	}

	if !p.expectPeek(token.RightParen) {
		return nil
	}

	return parameterTypes
}
