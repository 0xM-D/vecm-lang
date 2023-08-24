package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/0xM-D/interpreter/ast"
)

func (p *Parser) parseFloat32Literal() ast.Expression {
	lit := &ast.Float32Literal{Token: p.curToken}

	value, err := strconv.ParseFloat(strings.TrimSuffix(p.curToken.Literal, "f"), 32)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float32", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = float32(value)

	return lit
}

func (p *Parser) parseFloat64Literal() ast.Expression {
	lit := &ast.Float64Literal{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float64", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}
