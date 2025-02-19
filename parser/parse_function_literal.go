package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{
		Token:      p.curToken,
		Type:       ast.FunctionType{Token: p.curToken, ParameterTypes: nil, ReturnType: nil},
		Parameters: nil,
		IsVariadic: false,
		Body:       nil,
	}

	if !p.expectPeek(token.LeftParen) {
		return nil
	}

	lit.Parameters, lit.Type.ParameterTypes, lit.IsVariadic = p.parseFunctionParameters()
	if lit.Parameters == nil || lit.Type.ParameterTypes == nil {
		return nil
	}

	if !p.expectPeek(token.DashArrow) {
		return nil
	}
	p.nextToken()

	lit.Type.ReturnType = p.parseType()
	if lit.Type.ReturnType == nil {
		return nil
	}

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, []ast.Type, bool) {
	identifiers := []*ast.Identifier{}
	types := []ast.Type{}
	isVariadic := false

	if p.peekTokenIs(token.RightParen) {
		p.nextToken()
		return identifiers, types, false
	}

	ident, paramType := p.parseFunctionParameter()
	if ident == nil || paramType == nil {
		return nil, nil, false
	}
	identifiers = append(identifiers, ident)
	types = append(types, paramType)

	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		if p.peekTokenIs(token.Unpack) {
			p.nextToken()
			isVariadic = true
			break
		}
		ident, paramType = p.parseFunctionParameter()
		if ident == nil || paramType == nil {
			return nil, nil, false
		}
		identifiers = append(identifiers, ident)
		types = append(types, paramType)
	}

	if !p.expectPeek(token.RightParen) {
		if isVariadic {
			p.newErrorf(nil, "Variadic function parameters must be the last parameter")
		}
		return nil, nil, isVariadic
	}

	return identifiers, types, isVariadic
}

func (p *Parser) parseFunctionParameter() (*ast.Identifier, ast.Type) {
	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.Colon) {
		return nil, nil
	}

	p.nextToken()
	paramType := p.parseType()

	return ident, paramType
}
