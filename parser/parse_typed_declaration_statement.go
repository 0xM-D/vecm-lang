package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseTypedDeclarationStatement(stmtType ast.Type) *ast.TypedDeclarationStatement {
	typedDeclarationStatementToken := p.curToken

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
		ident := p.parseIdentifier()

		declStmtToken := p.curToken
		p.nextToken() // consume semicolon

		return &ast.TypedDeclarationStatement{
			Token: typedDeclarationStatementToken,
			DeclarationStatement: ast.DeclarationStatement{
				Token:      declStmtToken,
				Name:       ident,
				IsConstant: stmtIsConstant,
				Type:       stmtType,
				Value:      nil,
			},
		}
	}

	declStmt := p.parseDeclarationStatement(token.Assign)

	if declStmt == nil {
		return nil
	}

	return &ast.TypedDeclarationStatement{
		Token: typedDeclarationStatementToken,
		DeclarationStatement: ast.DeclarationStatement{
			Token:      declStmt.Token,
			Name:       declStmt.Name,
			Value:      declStmt.Value,
			IsConstant: stmtIsConstant,
			Type:       stmtType,
		},
	}
}
