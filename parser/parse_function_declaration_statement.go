package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseFunctionDeclarationStatement() *ast.FunctionDeclarationStatement {
	stmt := &ast.FunctionDeclarationStatement{Token: p.curToken}

	p.nextToken() // "fn"

	stmt.Name = p.parseIdentifier().(*ast.Identifier)

	if !p.expectPeek(token.LeftParen) {
		return nil
	}

	stmt.Parameters, stmt.Type.ParameterTypes = p.parseFunctionParameters()

	if !p.expectPeek(token.DashArrow) {
		return nil
	}
	p.nextToken()

	stmt.Type.ReturnType = p.parseType()
	if stmt.Type.ReturnType == nil {
		return nil
	}

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}
