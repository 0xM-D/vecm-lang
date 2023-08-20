package lexer

import (
	"github.com/0xM-D/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_ASSIGN, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		tok = newToken(token.PLUS, l.ch)
		switch l.peekChar() {
		case '=':
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_ASSIGN, Literal: string(ch) + string(l.ch)}
		case '>':
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DASH_ARROW, Literal: string(ch) + string(l.ch)}
		default:
			tok = newToken(token.MINUS, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.SLASH_ASSIGN, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ASTERISK_ASSIGN, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASTERISK, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		if l.peekChar() == ']' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ARRAY_TYPE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LBRACKET, l.ch)
		}
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '.':
		if isDigit(l.peekChar()) {
			tok.Type, tok.Literal = l.readNumber()
		} else {
			tok = newToken(token.ACCESS, l.ch)
		}
	case ':':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DECL_ASSIGN, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.COLON, l.ch)
		}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type, tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}
