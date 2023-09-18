package lexer

import (
	"github.com/0xM-D/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	linen        int
	coln         int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaceAndComments()

	switch l.ch {
	case '=':
		tok = l.getTokenWithPeek(token.ASSIGN, TokenMapping{'=', token.EQ})
	case ';':
		tok = l.newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = l.newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = l.newToken(token.RPAREN, string(l.ch))
	case ',':
		tok = l.newToken(token.COMMA, string(l.ch))
	case '+':
		tok = l.getTokenWithPeek(token.PLUS, TokenMapping{'=', token.PLUS_ASSIGN})
	case '-':
		tok = l.newToken(token.PLUS, string(l.ch))
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
			tok = l.newToken(token.MINUS, string(l.ch))
		}
	case '!':
		tok = l.getTokenWithPeek(token.BANG, TokenMapping{'=', token.NOT_EQ})
	case '/':
		tok = l.getTokenWithPeek(token.SLASH, TokenMapping{'=', token.SLASH_ASSIGN})
	case '*':
		tok = l.getTokenWithPeek(token.ASTERISK, TokenMapping{'=', token.ASTERISK_ASSIGN})
	case '<':
		tok = l.getTokenWithPeek(token.LT, TokenMapping{'=', token.LTE}, TokenMapping{'<', token.B_SHIFT_L})
	case '>':
		tok = l.getTokenWithPeek(token.GT, TokenMapping{'=', token.GTE}, TokenMapping{'>', token.B_SHIFT_R})
	case '{':
		tok = l.newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = l.newToken(token.RBRACE, string(l.ch))
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = l.getTokenWithPeek(token.LBRACKET, TokenMapping{']', token.ARRAY_TYPE})
	case ']':
		tok = l.newToken(token.RBRACKET, string(l.ch))
	case '.':
		if isDigit(l.peekChar()) {
			tok.Type, tok.Literal = l.readNumber()
		} else {
			tok = l.newToken(token.ACCESS, string(l.ch))
		}
	case ':':
		tok = l.getTokenWithPeek(token.COLON, TokenMapping{'=', token.DECL_ASSIGN})
	case '&':
		tok = l.getTokenWithPeek(token.B_AND, TokenMapping{'&', token.AND})
	case '|':
		tok = l.getTokenWithPeek(token.B_OR, TokenMapping{'|', token.OR})
	case '^':
		tok = l.newToken(token.B_XOR, string(l.ch))
	case '~':
		tok = l.newToken(token.B_INV, string(l.ch))
	case '?':
		tok = l.newToken(token.TERNARY_IF, string(l.ch))
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type, tok.Literal = l.readNumber()
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.ch))
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}
