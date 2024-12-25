package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseImportType() *ast.ImportType {
	switch p.curToken.Type {
	case token.CImportType:
		fallthrough
	case token.LLVMImportType:
		importTypeToken := p.curToken
		importType := &ast.ImportType{Token: importTypeToken, Value: p.curToken.Literal}

		p.nextToken()
		return importType
	default:
		return nil
	}
}
