package parser

import (
	"github.com/0xM-D/interpreter/token"
)

const (
	_ int = iota
	LOWEST
	COLON         // a: b pair expressions or ternary else
	ASSIGN        // a = b
	TERNARY_IF    // a ? b : c
	BITWISE_OR    // |, ^
	BITWISE_AND   // &
	BOOLEAN_OR    // ||
	BOOLEAN_AND   // &&
	EQUALS        // ==
	BITWISE_SHIFT // << or >>
	LESSGREATER   // > or <
	SUM           // +
	PRODUCT       // *
	PREFIX        // -X or !X or ~X
	CALL          // myFunction(X)
	INDEX         // array[index]
	ACCESS        // type.member or type.memberFn()
	TYPECAST      // 3 as uint8
)

var precedences = map[token.TokenType]int{
	token.B_INV:           PREFIX,
	token.B_AND:           BITWISE_AND,
	token.B_OR:            BITWISE_OR,
	token.B_XOR:           BITWISE_OR,
	token.AND:             BOOLEAN_AND,
	token.OR:              BOOLEAN_OR,
	token.EQ:              EQUALS,
	token.NOT_EQ:          EQUALS,
	token.B_SHIFT_L:       BITWISE_SHIFT,
	token.B_SHIFT_R:       BITWISE_SHIFT,
	token.LT:              LESSGREATER,
	token.GT:              LESSGREATER,
	token.LTE:             LESSGREATER,
	token.GTE:             LESSGREATER,
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
	token.QUESTIONMARK:    TERNARY_IF,
	token.COLON:           COLON,
	token.AS:              TYPECAST,
}
