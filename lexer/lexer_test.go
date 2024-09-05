package lexer_test

import (
	"testing"

	"github.com/DustTheory/interpreter/lexer"
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
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LeftParen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.LeftParen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.RightParen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.Gt, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.LeftParen, "("},
		{token.Int, "5"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Int, "10"},
		{token.Eq, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.NotEq, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		{token.String, "foobar"},
		{token.String, "foo bar"},
		{token.LeftBracket, "["},
		{token.Int, "1"},
		{token.Comma, ","},
		{token.Int, "2"},
		{token.RightBracket, "]"},
		{token.Semicolon, ";"},
		{token.LeftBrace, "{"},
		{token.String, "foo"},
		{token.Colon, ":"},
		{token.String, "bar"},
		{token.RightBrace, "}"},
		{token.Ident, "a"},
		{token.Assign, "="},
		{token.Int, "3"},
		{token.Ident, "str"},
		{token.PlusAssign, "+="},
		{token.String, "f"},
		{token.Ident, "a"},
		{token.MinusAssign, "-="},
		{token.Int, "4"},
		{token.Ident, "bff"},
		{token.AsteriskAssing, "*="},
		{token.String, "nadza"},
		{token.Ident, "div"},
		{token.SlashAssign, "/="},
		{token.Int, "1"},
		{token.Int, "5"},
		{token.Lte, "<="},
		{token.Int, "10"},
		{token.Gte, ">="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.True, "true"},
		{token.And, "&&"},
		{token.False, "false"},
		{token.Or, "||"},
		{token.True, "true"},
		{token.BitwiseInv, "~"},
		{token.Int, "0"},
		{token.BitwiseXor, "^"},
		{token.LeftParen, "("},
		{token.Int, "123"},
		{token.BitwiseAnd, "&"},
		{token.Int, "111"},
		{token.BitwiseOr, "|"},
		{token.Int, "0"},
		{token.RightParen, ")"},
		{token.Ident, "foo"},
		{token.BitwiseShiftL, "<<"},
		{token.Int, "3"},
		{token.Ident, "foo"},
		{token.BitwiseShiftR, ">>"},
		{token.Ident, "bar"},
		{token.Float32, "1.2f"},
		{token.Float32, ".123434f"},
		{token.Float32, "3f"},
		{token.Float64, "1.2312"},
		{token.Float64, ".1"},
		{token.Float64, "123.1"},
		{token.Int, "123"},
		{token.For, "for"},
		{token.LeftParen, "("},
		{token.Semicolon, ";"},
		{token.Semicolon, ";"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.RightBrace, "}"},
		{token.For, "for"},
		{token.LeftParen, "("},
		{token.Ident, "int"},
		{token.Ident, "i"},
		{token.Assign, "="},
		{token.Int, "0"},
		{token.Semicolon, ";"},
		{token.Ident, "i"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Ident, "i"},
		{token.PlusAssign, "+="},
		{token.Int, "1"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Const, "const"},
		{token.Ident, "x"},
		{token.Assign, "="},
		{token.Ident, "i"},
		{token.Asterisk, "*"},
		{token.Ident, "i"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Int, "123"},
		{token.For, "for"},
		{token.LeftParen, "("},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Ident, "condition"},
		{token.Questionmark, "?"},
		{token.True, "true"},
		{token.Colon, ":"},
		{token.False, "false"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

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
