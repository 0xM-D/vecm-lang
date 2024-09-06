package parser

import (
	"math/big"

	"github.com/DustTheory/interpreter/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	integerToken := p.curToken

	//nolint:mnd // 10 is the base for parsing the integer
	value, ok := new(big.Int).SetString(p.curToken.Literal, 10)

	lit := &ast.IntegerLiteral{
		Token: integerToken,
		Value: *value,
	}

	if !ok {
		p.newErrorf(lit, "could not parse %q as integer", p.curToken.Literal)
		return nil
	}

	return lit
}
