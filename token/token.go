package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Linen   int
	Coln    int
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals.
	IDENT   TokenType = "IDENT"
	INT     TokenType = "INT"
	FLOAT32 TokenType = "FLOAT32"
	FLOAT64 TokenType = "FLOAT64"
	STRING  TokenType = "STRING"

	// Operators.
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	ACCESS   TokenType = "."

	// Assignment operators.
	ASSIGN          TokenType = "="
	DECL_ASSIGN     TokenType = ":="
	PLUS_ASSIGN     TokenType = "+="
	MINUS_ASSIGN    TokenType = "-="
	ASTERISK_ASSIGN TokenType = "*="
	SLASH_ASSIGN    TokenType = "/="

	// Comparison operators.
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="
	EQ     = "=="
	NOT_EQ = "!="

	// Boolean operators.
	AND = "&&"
	OR  = "||"

	// Bitwise operators.
	B_AND     = "&"
	B_OR      = "|"
	B_XOR     = "^"
	B_INV     = "~"
	B_SHIFT_L = "<<"
	B_SHIFT_R = ">>"

	// Delimiters.
	COMMA        TokenType = ","
	SEMICOLON    TokenType = ";"
	COLON        TokenType = ":"
	QUESTIONMARK TokenType = "?"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// Keywords.
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	CONST    TokenType = "CONST"
	FOR      TokenType = "FOR"
	AS       TokenType = "AS"

	// Type tokens.
	DASH_ARROW    TokenType = "->"
	EQUALS_ARROW  TokenType = "=>"
	NEW           TokenType = "new"
	MAP_TYPE      TokenType = "map"
	ARRAY_TYPE    TokenType = "[]"
	FUNCTION_TYPE TokenType = "function"

	// import / export.
	IMPORT TokenType = "import"
	FROM   TokenType = "from"
	EXPORT TokenType = "export"
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
	"for":      FOR,
	"as":       AS,
	"new":      NEW,
	"import":   IMPORT,
	"from":     FROM,
	"export":   EXPORT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
