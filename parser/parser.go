package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},

		curToken: token.Token{
			Type:    token.EOF,
			Literal: "",
			Linen:   0,
			Coln:    0,
		},
		peekToken: token.Token{
			Type:    token.EOF,
			Literal: "",
			Linen:   0,
			Coln:    0,
		},

		prefixParseFns: make(map[token.Type]prefixParseFn),
		infixParseFns:  make(map[token.Type]infixParseFn),
	}

	p.prefixParseFns = make(map[token.Type]prefixParseFn)

	p.registerPrefix(token.Ident, p.parseIdentifierAsExpression)
	p.registerPrefix(token.Int, p.parseIntegerLiteral)
	p.registerPrefix(token.Float32, p.parseFloat32Literal)
	p.registerPrefix(token.Float64, p.parseFloat64Literal)
	p.registerPrefix(token.True, p.parseBooleanLiteral)
	p.registerPrefix(token.False, p.parseBooleanLiteral)
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.BitwiseInv, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.LeftParen, p.parseGroupedExpression)
	p.registerPrefix(token.Function, p.parseFunctionLiteral)
	p.registerPrefix(token.String, p.parseStringLiteral)
	p.registerPrefix(token.New, p.parseNewExpression)

	p.infixParseFns = make(map[token.Type]infixParseFn)
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Slash, p.parseInfixExpression)
	p.registerInfix(token.Asterisk, p.parseInfixExpression)
	p.registerInfix(token.Eq, p.parseInfixExpression)
	p.registerInfix(token.NotEq, p.parseInfixExpression)
	p.registerInfix(token.Lt, p.parseInfixExpression)
	p.registerInfix(token.Gt, p.parseInfixExpression)
	p.registerInfix(token.Lte, p.parseInfixExpression)
	p.registerInfix(token.Gte, p.parseInfixExpression)
	p.registerInfix(token.And, p.parseInfixExpression)
	p.registerInfix(token.Or, p.parseInfixExpression)
	p.registerInfix(token.BitwiseAnd, p.parseInfixExpression)
	p.registerInfix(token.BitwiseOr, p.parseInfixExpression)
	p.registerInfix(token.BitwiseXor, p.parseInfixExpression)
	p.registerInfix(token.BitwiseShiftL, p.parseInfixExpression)
	p.registerInfix(token.BitwiseShiftR, p.parseInfixExpression)
	p.registerInfix(token.LeftParen, p.parseCallExpression)
	p.registerInfix(token.LeftBracket, p.parseIndexExpression)
	p.registerInfix(token.Assign, p.parseInfixExpression)
	p.registerInfix(token.PlusAssign, p.parseInfixExpression)
	p.registerInfix(token.MinusAssign, p.parseInfixExpression)
	p.registerInfix(token.AsteriskAssing, p.parseInfixExpression)
	p.registerInfix(token.SlashAssign, p.parseInfixExpression)
	p.registerInfix(token.Access, p.parseAccessExpression)
	p.registerInfix(token.Questionmark, p.parseTernaryExpression)
	p.registerInfix(token.As, p.parseExplicitTypeCast)
	p.registerInfix(token.Colon, p.prasePairExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) newErrorf(node ast.Node, format string, a ...interface{}) {
	var linen, coln int
	if node == nil {
		linen, coln = p.l.GetLocation()
	} else {
		linen = node.TokenValue().Linen
		coln = node.TokenValue().Coln
	}

	p.errors = append(p.errors, lexer.NewError(linen, coln, p.getLine(linen), format, a...))
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
