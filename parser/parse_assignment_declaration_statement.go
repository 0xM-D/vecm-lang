package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseAssignmentDeclarationStatement() *ast.AssignmentDeclarationStatement {
	declStmt := p.parseDeclarationStatement(token.DECL_ASSIGN)
	if declStmt == nil {
		return nil
	}
	return &ast.AssignmentDeclarationStatement{DeclarationStatement: *declStmt}
}
