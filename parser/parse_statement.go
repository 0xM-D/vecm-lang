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
		fallthrough
	case token.MAP_TYPE:
		fallthrough
	case token.ARRAY_TYPE:
		fallthrough
	case token.FUNCTION_TYPE:
		return p.parseTypedDeclarationStatement(nil)
	case token.FOR:
		return p.parseForStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.IMPORT:
		return p.parseImportStatement()
	case token.EXPORT:
		return p.parseExportStatement()
	case token.FUNCTION:
		switch p.peekToken.Type {
		case token.IDENT:
			return p.parseFunctionDeclarationStatement()
		}
		fallthrough
	case token.IDENT:
		switch p.peekToken.Type {
		case token.DECL_ASSIGN:
			return p.parseAssignmentDeclarationStatement()
		case token.IDENT:
			fallthrough
		case token.ARRAY_TYPE:
			fallthrough
		case token.LBRACE:
			return p.parseTypedDeclarationStatement(nil)
		}
		fallthrough
	default:
		return p.parseExpressionStatement()
	}
}
