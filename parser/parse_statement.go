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
	case token.MAP_TYPE:
		return p.parseTypedDeclarationStatement()
	case token.IDENT:
		switch p.peekToken.Type {
		case token.DECL_ASSIGN:
			return p.parseAssignmentDeclarationStatement()
		case token.IDENT:
			fallthrough
		case token.ARRAY_TYPE:
			fallthrough
		case token.LBRACE:
			return p.parseTypedDeclarationStatement()
		}
		fallthrough
	default:
		return p.parseExpressionStatement()
	}
}
