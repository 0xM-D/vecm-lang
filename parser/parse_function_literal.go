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
	if lit.Parameters == nil || lit.Type.ParameterTypes == nil {
		return nil
	}

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

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, []*ast.FunctionParameterType) {
	identifiers := []*ast.Identifier{}
	types := []*ast.FunctionParameterType{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers, types
	}

	ident, param_type := p.parseFunctionParameter()
	if ident == nil || param_type == nil {
		return nil, nil
	}
	identifiers = append(identifiers, ident)
	types = append(types, param_type)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		ident, param_type := p.parseFunctionParameter()
		if ident == nil || param_type == nil {
			return nil, nil
		}
		identifiers = append(identifiers, ident)
		types = append(types, param_type)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, nil
	}

	return identifiers, types
}

// func (p *Parser) parseFunctionParameter() (*ast.Identifier, *ast.FunctionParameterType) {
// 	p.nextToken()
// 	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
// 	param_type := &ast.FunctionParameterType{}

// 	if p.peekTokenIs(token.COLON) {
// 		param_type.IsOptional = false
// 		p.nextToken()
// 	} else if p.peekTokenIs(token.QUESTIONMARK_COLON) {
// 		param_type.IsOptional = true
// 		p.nextToken()
// 	} else {
// 		p.newError(nil, "expected : or ?: token after function parameter identifier, got %s instead", p.peekToken.Literal)
// 		return nil, nil
// 	}
// 	p.nextToken()

// 	param_type.Type = p.parseType()

// 	return ident, param_type
// }

func (p *Parser) parseFunctionParameter() (*ast.Identifier, *ast.FunctionParameterType) {
	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.COLON) {
		return nil, nil
	}

	p.nextToken()
	param_type := p.parseType()

	return ident, &ast.FunctionParameterType{Type: param_type, IsOptional: false}
}
