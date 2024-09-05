package lexer

import (
	"github.com/DustTheory/interpreter/token"
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
	l := &Lexer{input: input, readPosition: 0, linen: 1, coln: 0, position: 0, ch: 0}
	l.readChar()
	return l
}

//nolint:funlen,gocyclo,cyclop // This large switch statement is quite nifty
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaceAndComments()

	switch l.ch {
	case '=':
		tok = l.getTokenWithPeek(token.Assign, TokenMapping{'=', token.Eq})
	case ';':
		tok = l.newToken(token.Semicolon, string(l.ch))
	case '(':
		tok = l.newToken(token.LeftParen, string(l.ch))
	case ')':
		tok = l.newToken(token.RightParen, string(l.ch))
	case ',':
		tok = l.newToken(token.Comma, string(l.ch))
	case '+':
		tok = l.getTokenWithPeek(token.Plus, TokenMapping{'=', token.PlusAssign})
	case '-':
		tok = l.newToken(token.Plus, string(l.ch))
		switch l.peekChar() {
		case '=':
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MinusAssign, Literal: string(ch) + string(l.ch), Linen: l.linen, Coln: l.coln}
		case '>':
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DashArrow, Literal: string(ch) + string(l.ch), Linen: l.linen, Coln: l.coln}
		default:
			tok = l.newToken(token.Minus, string(l.ch))
		}
	case '!':
		tok = l.getTokenWithPeek(token.Bang, TokenMapping{'=', token.NotEq})
	case '/':
		tok = l.getTokenWithPeek(token.Slash, TokenMapping{'=', token.SlashAssign})
	case '*':
		tok = l.getTokenWithPeek(token.Asterisk, TokenMapping{'=', token.AsteriskAssing})
	case '<':
		tok = l.getTokenWithPeek(token.Lt, TokenMapping{'=', token.Lte}, TokenMapping{'<', token.BitwiseShiftL})
	case '>':
		tok = l.getTokenWithPeek(token.Gt, TokenMapping{'=', token.Gte}, TokenMapping{'>', token.BitwiseShiftR})
	case '{':
		tok = l.newToken(token.LeftBrace, string(l.ch))
	case '}':
		tok = l.newToken(token.RightBrace, string(l.ch))
	case '"':
		tok.Type = token.String
		tok.Literal = l.readString()
	case '[':
		tok = l.getTokenWithPeek(token.LeftBracket, TokenMapping{']', token.ArrayType})
	case ']':
		tok = l.newToken(token.RightBracket, string(l.ch))
	case '.':
		if isDigit(l.peekChar()) {
			tok.Type, tok.Literal = l.readNumber()
		} else {
			tok = l.newToken(token.Access, string(l.ch))
		}
	case ':':
		tok = l.getTokenWithPeek(token.Colon, TokenMapping{'=', token.DeclAssign})
	case '&':
		tok = l.getTokenWithPeek(token.BitwiseAnd, TokenMapping{'&', token.And})
	case '|':
		tok = l.getTokenWithPeek(token.BitwiseOr, TokenMapping{'|', token.Or})
	case '^':
		tok = l.newToken(token.BitwiseXor, string(l.ch))
	case '~':
		tok = l.newToken(token.BitwiseInv, string(l.ch))
	case '?':
		tok = l.newToken(token.Questionmark, string(l.ch))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		switch {
		case isLetter(l.ch):
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		case isDigit(l.ch):
			tok.Type, tok.Literal = l.readNumber()
			return tok
		default:
			tok = l.newToken(token.Illegal, string(l.ch))
		}
	}
	l.readChar()
	return tok
}
