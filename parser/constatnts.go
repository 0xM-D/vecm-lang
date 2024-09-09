package parser

import (
	"github.com/DustTheory/interpreter/token"
)

const (
	_ int = iota
	Lowest
	Colon        // a: b pair expressions or ternary else
	Assign       // a = b
	TernaryIf    // a ? b : c
	BitwiseOr    // |, ^
	BitwiseAnd   // &
	BooleanOr    // ||
	BooleanAnd   // &&
	Equals       // ==
	BitwiseShift // << or >>
	Comparison   // > or <
	Sum          // +
	Product      // *
	Prefix       // -X or !X or ~X
	Call         // myFunction(X)
	Index        // array[index]
	Access       // type.member or type.memberFn()
	Typecast     // 3 as uint8
)

//
var precedences = map[token.Type]int{
	token.BitwiseInv:     Prefix,
	token.BitwiseAnd:     BitwiseAnd,
	token.BitwiseOr:      BitwiseOr,
	token.BitwiseXor:     BitwiseOr,
	token.And:            BooleanAnd,
	token.Or:             BooleanOr,
	token.Eq:             Equals,
	token.NotEq:          Equals,
	token.BitwiseShiftL:  BitwiseShift,
	token.BitwiseShiftR:  BitwiseShift,
	token.Lt:             Comparison,
	token.Gt:             Comparison,
	token.Lte:            Comparison,
	token.Gte:            Comparison,
	token.Plus:           Sum,
	token.Minus:          Sum,
	token.Slash:          Product,
	token.Asterisk:       Product,
	token.LeftParen:      Call,
	token.LeftBracket:    Index,
	token.Assign:         Assign,
	token.DeclAssign:     Assign,
	token.PlusAssign:     Assign,
	token.MinusAssign:    Assign,
	token.AsteriskAssing: Assign,
	token.SlashAssign:    Assign,
	token.Access:         Access,
	token.Questionmark:   TernaryIf,
	token.Colon:          Colon,
	token.As:             Typecast,
}
