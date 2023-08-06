package ast

import (
	"github.com/0xM-D/interpreter/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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
	ReturnValue Expression
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
	Value int64
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

type Boolean struct {
	Token token.Token
	Value bool
}

type IfExpression struct {
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
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

type TypedDeclarationStatement struct {
	DeclarationStatement
}

type AssignmentDeclarationStatement struct {
	DeclarationStatement
}

type VariableUpdateStatement struct {
	Token    token.Token
	Left     *Identifier
	Operator string
	Right    Expression
}

type NamedType struct {
	Token    token.Token
	TypeName Identifier
}

type ReturnType struct {
	Type *NamedType // Nullable
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
	ReturnType     ReturnType
}
