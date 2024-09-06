package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	importStatementToken := p.curToken

	importIdentifiers, importAll := p.parseImportIdentifiers()

	if importIdentifiers == nil && !importAll {
		return nil
	}

	p.nextToken() // "from"

	importPathString := p.parseStringLiteral()
	if importPathString == nil {
		return nil
	}

	importPath := importPathString.(*ast.StringLiteral).Value

	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return &ast.ImportStatement{
		Token:               importStatementToken,
		ImportPath:          importPath,
		ImportAll:           importAll,
		ImportedIdentifiers: importIdentifiers,
	}
}

func (p *Parser) parseImportIdentifiers() ([]*ast.Identifier, bool) {
	importAll := p.peekTokenIs(token.Asterisk)

	if importAll {
		p.nextToken() // Swallow "import"
		p.nextToken() // Swallow "*"
		return nil, true
	}

	importIdentifiers := p.parseImportLst()
	return importIdentifiers, false
}

func (p *Parser) parseImportLst() []*ast.Identifier {
	list := []*ast.Identifier{}
	if p.peekTokenIs(token.From) {
		p.nextToken()
		return list
	}
	p.nextToken()
	if !p.curTokenIs(token.Ident) {
		p.newErrorf(nil, "Expected identifier in import statement. got=%T", p.curToken)
		return nil
	}
	list = append(list, p.parseIdentifier())
	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		if !p.curTokenIs(token.Ident) {
			p.newErrorf(nil, "Expected identifier in import statement. got=%T", p.curToken)
			return nil
		}
		list = append(list, p.parseIdentifier())
	}
	if !p.expectPeek(token.From) {
		return nil
	}
	return list
}
