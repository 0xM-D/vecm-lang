package parser

import "github.com/DustTheory/interpreter/ast"

func (p *Parser) parseExplicitTypeCast(left ast.Expression) ast.Expression {
	expr := &ast.TypeCastExpression{Token: p.curToken, Left: left}

	p.nextToken() // as

	expr.Type = p.parseType()

	if expr.Type == nil {
		return nil
	}

	return expr
}
