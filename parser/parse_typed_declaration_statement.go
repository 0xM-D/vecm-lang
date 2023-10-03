package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseTypedDeclarationStatement(stmtType ast.Type) *ast.TypedDeclarationStatement {
	stmt := &ast.TypedDeclarationStatement{Token: p.curToken}

	var stmtIsConstant bool

	if p.curTokenIs(token.CONST) {
		stmtIsConstant = true
		p.nextToken()
	}

	if !p.peekTokenIs(token.ASSIGN) && stmtType == nil {
		stmtType = p.parseType()
		p.nextToken()
	}

	declStmt := p.parseDeclarationStatement(token.ASSIGN)

	if declStmt == nil {
		return nil
	}

	stmt.DeclarationStatement = *declStmt
	stmt.DeclarationStatement.IsConstant = stmtIsConstant
	stmt.DeclarationStatement.Type = stmtType

	return stmt
}
