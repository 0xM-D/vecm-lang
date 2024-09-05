package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	case token.Const:
		fallthrough
	case token.MapType:
		fallthrough
	case token.ArrayType:
		fallthrough
	case token.FunctionType:
		return p.parseTypedDeclarationStatement(nil)
	case token.For:
		return p.parseForStatement()
	case token.If:
		return p.parseIfStatement()
	case token.Import:
		return p.parseImportStatement()
	case token.Export:
		return p.parseExportStatement()
	case token.Function:
		switch p.peekToken.Type {
		case token.Ident:
			return p.parseFunctionDeclarationStatement()
		}
		fallthrough
	case token.Ident:
		switch p.peekToken.Type {
		case token.DeclAssign:
			return p.parseAssignmentDeclarationStatement()
		case token.Ident:
			fallthrough
		case token.ArrayType:
			fallthrough
		case token.LeftBrace:
			return p.parseTypedDeclarationStatement(nil)
		}
		fallthrough
	default:
		return p.parseExpressionStatement()
	}
}
