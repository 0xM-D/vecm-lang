package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	importStatement := &ast.ImportStatement{Token: p.curToken}

	importStatement.ImportAll = p.peekTokenIs(token.ASTERISK)

	if importStatement.ImportAll {
		p.nextToken()
		p.nextToken()
	} else {
		importStatement.ImportedIdentifiers = p.parseImportLst()
		if importStatement.ImportedIdentifiers == nil {
			return nil
		}
	}

	p.nextToken() // "from"

	importPathString := p.parseStringLiteral()
	if importPathString == nil {
		return nil
	}
	importStatement.ImportPath = importPathString.(*ast.StringLiteral).Value

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return importStatement
}

func (p *Parser) parseImportLst() []*ast.Identifier {
	list := []*ast.Identifier{}
	if p.peekTokenIs(token.FROM) {
		p.nextToken()
		return list
	}
	p.nextToken()
	if !p.curTokenIs(token.IDENT) {
		p.newError(nil, "Expected identifier in import statement. got=%T", p.curToken)
		return nil
	}
	list = append(list, p.parseIdentifier().(*ast.Identifier))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		if !p.curTokenIs(token.IDENT) {
			p.newError(nil, "Expected identifier in import statement. got=%T", p.curToken)
			return nil
		}
		list = append(list, p.parseIdentifier().(*ast.Identifier))
	}
	if !p.expectPeek(token.FROM) {
		return nil
	}
	return list
}
