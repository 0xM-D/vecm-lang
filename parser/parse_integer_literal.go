package parser

import (
	"math/big"

	"github.com/DustTheory/interpreter/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, ok := new(big.Int).SetString(p.curToken.Literal, 10)
	if !ok {
		p.newError(lit, "could not parse %q as integer", p.curToken.Literal)
		return nil
	}

	lit.Value = *value

	return lit
}
