package lexer

import (
	"testing"

	"github.com/DustTheory/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
!-/ *5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
a = 3
str += "f"
a -= 4
bff *= "nadza"
div /= 1
5 <= 10 >= 5;
true && false || true
~0 ^ (123 & 111 | 0)
foo << 3
foo >> bar
1.2f
.123434f
3f
1.2312
.1
123.1
123
for(;;){}
for(int i = 0; i < 10; i+=1){
	const x = i * i;
}
// test
// test test
123 // test
for(//////
x + /* test */ y
/*
	multi
line
  comment */
condition ? true : false
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.IDENT, "str"},
		{token.PLUS_ASSIGN, "+="},
		{token.STRING, "f"},
		{token.IDENT, "a"},
		{token.MINUS_ASSIGN, "-="},
		{token.INT, "4"},
		{token.IDENT, "bff"},
		{token.ASTERISK_ASSIGN, "*="},
		{token.STRING, "nadza"},
		{token.IDENT, "div"},
		{token.SLASH_ASSIGN, "/="},
		{token.INT, "1"},
		{token.INT, "5"},
		{token.LTE, "<="},
		{token.INT, "10"},
		{token.GTE, ">="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.TRUE, "true"},
		{token.AND, "&&"},
		{token.FALSE, "false"},
		{token.OR, "||"},
		{token.TRUE, "true"},
		{token.B_INV, "~"},
		{token.INT, "0"},
		{token.B_XOR, "^"},
		{token.LPAREN, "("},
		{token.INT, "123"},
		{token.B_AND, "&"},
		{token.INT, "111"},
		{token.B_OR, "|"},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.IDENT, "foo"},
		{token.B_SHIFT_L, "<<"},
		{token.INT, "3"},
		{token.IDENT, "foo"},
		{token.B_SHIFT_R, ">>"},
		{token.IDENT, "bar"},
		{token.FLOAT32, "1.2f"},
		{token.FLOAT32, ".123434f"},
		{token.FLOAT32, "3f"},
		{token.FLOAT64, "1.2312"},
		{token.FLOAT64, ".1"},
		{token.FLOAT64, "123.1"},
		{token.INT, "123"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.SEMICOLON, ";"},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.IDENT, "int"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.PLUS_ASSIGN, "+="},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.CONST, "const"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.IDENT, "i"},
		{token.ASTERISK, "*"},
		{token.IDENT, "i"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "123"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.IDENT, "condition"},
		{token.QUESTIONMARK, "?"},
		{token.TRUE, "true"},
		{token.COLON, ":"},
		{token.FALSE, "false"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
