package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseType() ast.Type {

	var result ast.Type

	if p.curToken.Type == token.ARRAY_TYPE {
		p.nextToken()
		elementType := p.parseType()
		if elementType == nil {
			return nil
		}
		return ast.ArrayType{Token: p.curToken, ElementType: elementType}
	}

	switch p.curToken.Type {
	case token.MAP_TYPE:
		result = p.parseMapType()
	case token.FUNCTION_TYPE:
		result = p.parseFunctionType()
	case token.IDENT:
		typeIdentifier := p.parseIdentifier().(*ast.Identifier)
		result = ast.NamedType{Token: p.curToken, TypeName: *typeIdentifier}
	default:
		return nil
	}
	return result
}

func (p *Parser) parseMapType() ast.Type {
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

	return ast.HashType{Token: p.curToken, KeyType: keyType, ValueType: valueType}
}

func (p *Parser) parseFunctionType() ast.Type {

	functionType := ast.FunctionType{Token: p.curToken}

	functionType.ParameterTypes = p.parseFunctionTypeParameters()

	if functionType.ParameterTypes == nil {
		return nil
	}

	if !p.expectPeek(token.DASH_ARROW) {
		return nil
	}
	p.nextToken()

	functionType.ReturnType = p.parseType()

	if functionType.ReturnType == nil {
		return nil
	}

	return functionType
}

func (p *Parser) parseFunctionTypeParameters() []*ast.FunctionParameterType {
	parameterTypes := []*ast.FunctionParameterType{}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	for {
		if p.peekTokenIs(token.RPAREN) {
			break
		}
		p.nextToken()

		param_type := &ast.FunctionParameterType{IsOptional: false}

		param_type.Type = p.parseType()
		if param_type.Type == nil {
			return nil
		}

		if p.peekTokenIs(token.QUESTIONMARK) {
			param_type.IsOptional = true
			p.nextToken()
		}

		parameterTypes = append(parameterTypes, param_type)

		if !p.peekTokenIs(token.COMMA) {
			break
		}
		p.nextToken()
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return parameterTypes
}
