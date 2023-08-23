package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	ACCESS   = "."

	// Assignment operators
	ASSIGN          = "="
	DECL_ASSIGN     = ":="
	PLUS_ASSIGN     = "+="
	MINUS_ASSIGN    = "-="
	ASTERISK_ASSIGN = "*="
	SLASH_ASSIGN    = "/="

	// Comparison operators
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="
	EQ     = "=="
	NOT_EQ = "!="

	// Boolean operators
	AND = "&&"
	OR  = "||"

	// Bitwise operators
	B_AND     = "&"
	B_OR      = "|"
	B_XOR     = "^"
	B_INV     = "~"
	B_SHIFT_L = "<<"
	B_SHIFT_R = ">>"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	CONST    = "CONST"

	// Type tokens
	DASH_ARROW    = "->"
	EQUALS_ARROW  = "=>"
	MAP_TYPE      = "map"
	ARRAY_TYPE    = "[]"
	FUNCTION_TYPE = "function"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"const":    CONST,
	"map":      MAP_TYPE,
	"function": FUNCTION_TYPE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
