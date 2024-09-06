package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseIfStatement() ast.Statement {
	// Swallow "if" token
	ifStatementToken := p.curToken
	p.nextToken()

	condition := p.parseExpression(Lowest)
	consequence := p.parseIfStatementConsequence()
	alternative := p.parseIfStatementAlternative()

	return &ast.IfStatement{
		Token:       ifStatementToken,
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

func (p *Parser) parseIfStatementConsequence() *ast.BlockStatement {
	p.nextToken() // Swallow "{" token
	if !p.curTokenIs(token.LeftBrace) {
		return &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
	}
	return p.parseBlockStatement()
}

func (p *Parser) parseIfStatementAlternative() *ast.BlockStatement {
	if p.peekTokenIs(token.Else) {
		p.nextToken() // Swallow "else" token
		p.nextToken() // Swallow "{" token
		if !p.curTokenIs(token.LeftBrace) {
			return &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{p.parseStatement()}}
		}
		return p.parseBlockStatement()
	}
	return nil
}
