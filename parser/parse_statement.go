package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.CONST:
		return p.parseConstDeclarationStatement()
	case token.IDENT:
		switch p.peekToken.Type {
		case token.IDENT:
			return p.parseTypedDeclarationStatement()
		case token.DECL_ASSIGN:
			return p.parseAssignmentDeclarationStatement()
		}
		fallthrough
	default:
		return p.parseExpressionStatement()
	}
}
