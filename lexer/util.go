package lexer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/DustTheory/interpreter/token"
)

func (l *Lexer) readChar() {
	if l.ch == '\n' {
		l.linen++
		l.coln = 0
	} else {
		l.coln++
	}
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespaceAndComments() {
	for l.ch != 0 {
		switch {
		case l.whiteSpaceAhead():
			l.skipWhitespace()
		case l.commentAhead():
			l.skipComments()
		default:
			return
		}
	}
}

func (l *Lexer) skipWhitespace() {
	for l.whiteSpaceAhead() {
		l.readChar()
	}
}

func (l *Lexer) skipComments() {
	if !l.commentAhead() {
		return
	}
	switch l.peekChar() {
	case '/':
		l.skipLineComment()
	case '*':
		l.skipBlockComment()
	}
}

func (l *Lexer) skipLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.readChar()
}

func (l *Lexer) skipBlockComment() {
	// skip "/*"
	l.readChar()
	l.readChar()

	for (l.ch != '*' || l.peekChar() != '/') && l.ch != 0 {
		l.readChar()
	}

	// skip "*/"
	l.readChar()
	l.readChar()
}

func (l *Lexer) whiteSpaceAhead() bool {
	return l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'
}

func (l *Lexer) commentAhead() bool {
	peekTwo := string(l.ch) + string(l.peekChar())
	return peekTwo == "//" || peekTwo == "/*"
}

func (l *Lexer) newToken(tokenType token.Type, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal, Linen: l.linen, Coln: l.coln}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (token.Type, string) {
	position := l.position
	tokenType := token.Int

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' && isDigit(l.peekChar()) {
		tokenType = token.Float64
		l.readChar()

		for isDigit(l.ch) {
			l.readChar()
		}
	}

	if l.ch == 'f' {
		tokenType = token.Float32
		l.readChar()
	}
	return tokenType, l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) GetLocation() (int, int) {
	return l.linen, l.coln
}

func (l *Lexer) GetLine(linen int) string {
	return strings.Split(l.input, "\n")[linen]
}

type TokenMapping struct {
	byte
	token.Type
}

func (l *Lexer) getTokenWithPeek(defaultToken token.Type, tokenMappings ...TokenMapping) token.Token {
	for _, tokenMapping := range tokenMappings {
		if l.peekChar() == tokenMapping.byte {
			ch := l.ch
			l.readChar()
			return token.Token{Type: tokenMapping.Type, Literal: string(ch) + string(l.ch), Linen: l.linen, Coln: l.coln}
		}
	}

	return l.newToken(defaultToken, string(l.ch))
}

func NewError(linen int, coln int, line string, format string, a ...interface{}) string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("Parser error at line %d, column %d:\n", linen+1, coln+1))
	out.WriteString(line)
	out.WriteByte('\n')
	out.WriteString(getASCIIArrow(coln))
	out.WriteString(fmt.Sprintf(format, a...))
	out.WriteByte('\n')

	return out.String()
}

func getASCIIArrow(coln int) string {
	var out bytes.Buffer
	for i := 0; i < coln; i++ {
		out.WriteByte(' ')
	}
	out.WriteByte('^')
	out.WriteByte('\n')
	return out.String()
}
