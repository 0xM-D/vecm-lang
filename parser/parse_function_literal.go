package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken, Type: ast.FunctionType{Token: p.curToken}}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters, lit.Type.ParameterTypes = p.parseFunctionParameters()

	if !p.expectPeek(token.DASH_ARROW) {
		return nil
	}
	p.nextToken()

	lit.Type.ReturnType = p.parseType()
	if lit.Type.ReturnType == nil {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, []ast.Type) {
	identifiers := []*ast.Identifier{}
	types := []ast.Type{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers, types
	}

	ident, param_type := p.parseFunctionParameter()
	identifiers = append(identifiers, ident)
	types = append(types, param_type)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		ident, param_type := p.parseFunctionParameter()
		identifiers = append(identifiers, ident)
		types = append(types, param_type)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, nil
	}

	return identifiers, types
}

func (p *Parser) parseFunctionParameter() (*ast.Identifier, ast.Type) {
	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.COLON) {
		return nil, nil
	}

	p.nextToken()
	param_type := p.parseType()

	return ident, param_type
}
