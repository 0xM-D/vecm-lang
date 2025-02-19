package ast

import (
	"math/big"

	"github.com/DustTheory/interpreter/token"
)

type Node interface {
	TokenLiteral() string
	String() string
	TokenValue() token.Token
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Type interface {
	Node
	typeNode()
}

type Program struct {
	Statements []Statement
}

type DeclarationStatement struct {
	Token      token.Token
	IsConstant bool
	Name       *Identifier
	Type       Type
	Value      Expression
}

type LetStatement struct {
	DeclarationStatement
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression `exhaustruct:"optional"`
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type IntegerLiteral struct {
	Token token.Token
	Value big.Int
}

type Float32Literal struct {
	Token token.Token
	Value float32
}

type Float64Literal struct {
	Token token.Token
	Value float64
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
	IsVariadic bool
	Type       FunctionType
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
	Type
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

type AccessExpression struct {
	Token token.Token
	Left  Expression
	Right Expression
}

type HashLiteral struct {
	Token    token.Token
	Pairs    map[Expression]Expression
	HashType HashType
}

type TypedDeclarationStatement struct {
	Token token.Token
	DeclarationStatement
}

type AssignmentDeclarationStatement struct {
	DeclarationStatement
}

type VoidType struct {
	Token token.Token
}

type NamedType struct {
	Token    token.Token
	TypeName Identifier
}

type ArrayType struct {
	Token       token.Token
	ElementType Type
}

type HashType struct {
	Token     token.Token
	KeyType   Type
	ValueType Type
}

type FunctionType struct {
	Token          token.Token
	ParameterTypes []Type
	ReturnType     Type
}

type PointerType struct {
	Token       token.Token
	PointeeType Type
}

type ForStatement struct {
	Token          token.Token
	Initialization Statement
	Condition      Statement
	AfterThought   Statement
	Body           *BlockStatement
}

type TernaryExpression struct {
	Token        token.Token
	Condition    Expression
	ValueIfTrue  Expression
	ValueIfFalse Expression
}

type TypeCastExpression struct {
	Token token.Token
	Left  Expression
	Type  Type
}

type NewExpression struct {
	Token              token.Token
	Type               Type
	InitializationList []Expression
}

type PairExpression struct {
	Token token.Token
	Left  Expression
	Right Expression
}

type ImportStatement struct {
	Token               token.Token
	ImportPath          string
	ImportAll           bool
	ImportedIdentifiers []*Identifier
}

type ExportStatement struct {
	Token token.Token
	Statement
}

type FunctionDeclarationStatement struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	IsVariadic bool
	Body       *BlockStatement
	Type       FunctionType
}

type CLangStatement struct {
	Token     token.Token
	CLangCode string
}
