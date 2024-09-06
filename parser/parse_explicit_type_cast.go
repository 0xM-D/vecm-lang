package parser

import "github.com/DustTheory/interpreter/ast"

func (p *Parser) parseExplicitTypeCast(left ast.Expression) ast.Expression {
	expr := &ast.TypeCastExpression{Token: p.curToken, Left: left, Type: nil}

	p.nextToken() // Swallow "as" token

	expr.Type = p.parseType()

	if expr.Type == nil {
		return nil
	}

	return expr
}
