package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	//nolint:exhaustive // This is exhaustive
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
		if p.peekToken.Type == token.Ident {
			return p.parseFunctionDeclarationStatement()
		}
		fallthrough
	case token.Ident:
		//nolint:exhaustive // This is exhaustive
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
