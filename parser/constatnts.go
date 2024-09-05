package parser

import (
	"github.com/DustTheory/interpreter/token"
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

var precedences = map[token.Type]int{
	token.BitwiseInv:           PREFIX,
	token.BitwiseAnd:           BITWISE_AND,
	token.BitwiseOr:            BITWISE_OR,
	token.BitwiseXor:           BITWISE_OR,
	token.And:             BOOLEAN_AND,
	token.Or:              BOOLEAN_OR,
	token.Eq:              EQUALS,
	token.NotEq:          EQUALS,
	token.BitwiseShiftL:       BITWISE_SHIFT,
	token.BitwiseShiftR:       BITWISE_SHIFT,
	token.Lt:              LESSGREATER,
	token.Gt:              LESSGREATER,
	token.Lte:             LESSGREATER,
	token.Gte:             LESSGREATER,
	token.Plus:            SUM,
	token.Minus:           SUM,
	token.Slash:           PRODUCT,
	token.Asterisk:        PRODUCT,
	token.LeftParen:          CALL,
	token.LeftBracket:        INDEX,
	token.Assign:          ASSIGN,
	token.DeclAssign:     ASSIGN,
	token.PlusAssign:     ASSIGN,
	token.MinusAssign:    ASSIGN,
	token.AsteriskAssing: ASSIGN,
	token.SlashAssign:    ASSIGN,
	token.Access:          ACCESS,
	token.Questionmark:    TERNARY_IF,
	token.Colon:           COLON,
	token.As:              TYPECAST,
}
