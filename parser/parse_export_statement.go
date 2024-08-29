package parser

import "github.com/DustTheory/interpreter/ast"

func (p *Parser) parseExportStatement() ast.Statement {
	exportStmt := &ast.ExportStatement{Token: p.curToken}

	p.nextToken() // "export"
	exportStmt.Statement = p.parseStatement()
	if exportStmt.Statement == nil {
		return nil
	}

	return exportStmt
}
