package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseTypedDeclarationStatement(stmtType ast.Type) *ast.TypedDeclarationStatement {
	stmt := &ast.TypedDeclarationStatement{Token: p.curToken}

	var stmtIsConstant bool

	if p.curTokenIs(token.Const) {
		stmtIsConstant = true
		p.nextToken()
	}

	if !p.peekTokenIs(token.Assign) && stmtType == nil {
		stmtType = p.parseType()
		p.nextToken()
	}

	if p.peekTokenIs(token.Semicolon) {
		ident := p.parseIdentifier().(*ast.Identifier)

		stmt.DeclarationStatement = ast.DeclarationStatement{
			Token:      p.curToken,
			Name:       ident,
			IsConstant: stmtIsConstant,
			Type:       stmtType,
			Value:      nil,
		}

		p.nextToken() // consume semicolon
		return stmt
	}

	declStmt := p.parseDeclarationStatement(token.Assign)

	if declStmt == nil {
		return nil
	}

	stmt.DeclarationStatement = *declStmt
	stmt.DeclarationStatement.IsConstant = stmtIsConstant
	stmt.DeclarationStatement.Type = stmtType

	return stmt
}
