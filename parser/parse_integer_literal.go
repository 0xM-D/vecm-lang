package parser

import (
	"strconv"

	"github.com/0xM-D/interpreter/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.newError(lit, "could not parse %q as integer", p.curToken.Literal)
		return nil
	}

	lit.Value = value

	return lit
}
