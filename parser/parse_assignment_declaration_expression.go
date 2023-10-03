package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseAssignmentDeclarationStatement() *ast.AssignmentDeclarationStatement {
	declStmt := p.parseDeclarationStatement(token.DECL_ASSIGN)
	if declStmt == nil {
		return nil
	}
	return &ast.AssignmentDeclarationStatement{DeclarationStatement: *declStmt}
}
