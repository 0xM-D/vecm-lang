package parser

import (
	"github.com/0xM-D/interpreter/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	ACCESS      // type.member or type.memberFn()
	INDEX       // array[index]
	ASSIGN      // a = b
)

var precedences = map[token.TokenType]int{
	token.EQ:              EQUALS,
	token.NOT_EQ:          EQUALS,
	token.LT:              LESSGREATER,
	token.GT:              LESSGREATER,
	token.PLUS:            SUM,
	token.MINUS:           SUM,
	token.SLASH:           PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.LPAREN:          CALL,
	token.LBRACKET:        INDEX,
	token.ASSIGN:          ASSIGN,
	token.DECL_ASSIGN:     ASSIGN,
	token.PLUS_ASSIGN:     ASSIGN,
	token.MINUS_ASSIGN:    ASSIGN,
	token.ASTERISK_ASSIGN: ASSIGN,
	token.SLASH_ASSIGN:    ASSIGN,
	token.ACCESS:          ACCESS,
}
