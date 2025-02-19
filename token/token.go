package token

type Type string

type Token struct {
	Type    Type
	Literal string
	Linen   int
	Coln    int
}

const (
	Illegal Type = "ILLEGAL"
	EOF     Type = "EOF"

	// Identifiers + literals.
	Ident   Type = "IDENT"
	Int     Type = "INT"
	Float32 Type = "FLOAT32"
	Float64 Type = "FLOAT64"
	String  Type = "STRING"

	// Operators.
	Plus     Type = "+"
	Minus    Type = "-"
	Bang     Type = "!"
	Asterisk Type = "*"
	Slash    Type = "/"
	Access   Type = "."

	// Assignment operators.
	Assign         Type = "="
	DeclAssign     Type = ":="
	PlusAssign     Type = "+="
	MinusAssign    Type = "-="
	AsteriskAssing Type = "*="
	SlashAssign    Type = "/="

	// Comparison operators.
	Lt    = "<"
	Gt    = ">"
	Lte   = "<="
	Gte   = ">="
	Eq    = "=="
	NotEq = "!="

	// Boolean operators.
	And = "&&"
	Or  = "||"

	// Bitwise operators.
	BitwiseAnd    = "&"
	BitwiseOr     = "|"
	BitwiseXor    = "^"
	BitwiseInv    = "~"
	BitwiseShiftL = "<<"
	BitwiseShiftR = ">>"

	// Delimiters.
	Comma        Type = ","
	Semicolon    Type = ";"
	Colon        Type = ":"
	Questionmark Type = "?"

	// Unpack operator.
	Unpack Type = "..."

	LeftParen    Type = "("
	RightParen   Type = ")"
	LeftBrace    Type = "{"
	RightBrace   Type = "}"
	LeftBracket  Type = "["
	RightBracket Type = "]"

	// Keywords.
	Function Type = "FUNCTION"
	Let      Type = "LET"
	True     Type = "TRUE"
	False    Type = "FALSE"
	If       Type = "IF"
	Else     Type = "ELSE"
	Return   Type = "RETURN"
	Const    Type = "CONST"
	For      Type = "FOR"
	As       Type = "AS"

	// Type tokens.
	DashArrow    Type = "->"
	EqualsArrow  Type = "=>"
	New          Type = "new"
	MapType      Type = "map"
	ArrayType    Type = "[]"
	FunctionType Type = "function"

	// import / export.
	Import Type = "import"
	From   Type = "from"
	Export Type = "export"

	// CLang.
	CLang Type = "CLang"
)

// Mapping of keywords to their respective token types.
var keywords = map[string]Type{
	"fn":       Function,
	"let":      Let,
	"true":     True,
	"false":    False,
	"if":       If,
	"else":     Else,
	"return":   Return,
	"const":    Const,
	"map":      MapType,
	"function": FunctionType,
	"for":      For,
	"as":       As,
	"new":      New,
	"import":   Import,
	"from":     From,
	"export":   Export,
	"CLang":    CLang,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
