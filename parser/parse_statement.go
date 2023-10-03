package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {

	if p.curToken.Type == token.ARRAY_TYPE {
		startToken := p.curToken
		arrayType := p.parseType()
		if arrayType == nil {
			return nil
		}
		if p.peekTokenIs(token.LBRACE) {
			arrayLiteral := p.parseArrayLiteralRest(&ast.ArrayLiteral{Token: startToken, Type: arrayType})
			if arrayLiteral == nil {
				return nil
			}
			expr := p.parseExpressionRest(LOWEST, arrayLiteral)
			if expr == nil {
				return nil
			}
			exprStmt := &ast.ExpressionStatement{
				Token:      p.curToken,
				Expression: expr,
			}
			if p.peekTokenIs(token.SEMICOLON) {
				p.nextToken()
			}
			return exprStmt
		} else {
			p.nextToken()
			return p.parseTypedDeclarationStatement(arrayType)
		}
	}

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.CONST:
		return p.parseTypedDeclarationStatement(nil)
	case token.MAP_TYPE:
		return p.parseTypedDeclarationStatement(nil)
	case token.FUNCTION_TYPE:
		return p.parseTypedDeclarationStatement(nil)
	case token.FOR:
		return p.parseForStatement()
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
