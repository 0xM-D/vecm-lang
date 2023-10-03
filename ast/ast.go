package ast

import (
	"bytes"
	"sort"
	"strings"

	"github.com/0xM-D/interpreter/token"
)

func (p *Program) TokenLiteral() string {
	return p.TokenValue().Literal
}

func (p *Program) TokenValue() token.Token {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenValue()
	} else {
		return token.Token{Type: token.EOF, Literal: "", Linen: 0, Coln: 0}
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) statementNode()          {}
func (ls *LetStatement) TokenLiteral() string    { return ls.Token.Literal }
func (ls *LetStatement) TokenValue() token.Token { return ls.Token }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (i *Identifier) expressionNode()         {}
func (i *Identifier) TokenLiteral() string    { return i.Token.Literal }
func (i *Identifier) TokenValue() token.Token { return i.Token }
func (i *Identifier) String() string          { return i.Value }

func (rs *ReturnStatement) statementNode()          {}
func (rs *ReturnStatement) TokenLiteral() string    { return rs.Token.Literal }
func (rs *ReturnStatement) TokenValue() token.Token { return rs.Token }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) statementNode()          {}
func (es *ExpressionStatement) TokenLiteral() string    { return es.Token.Literal }
func (es *ExpressionStatement) TokenValue() token.Token { return es.Token }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (il *IntegerLiteral) expressionNode()         {}
func (il *IntegerLiteral) TokenLiteral() string    { return il.Token.Literal }
func (il *IntegerLiteral) TokenValue() token.Token { return il.Token }
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (fp *Float64Literal) expressionNode()         {}
func (fp *Float64Literal) TokenLiteral() string    { return fp.Token.Literal }
func (fp *Float64Literal) TokenValue() token.Token { return fp.Token }
func (fp *Float64Literal) String() string {
	return fp.Token.Literal
}

func (fp *Float32Literal) expressionNode()         {}
func (fp *Float32Literal) TokenLiteral() string    { return fp.Token.Literal }
func (fp *Float32Literal) TokenValue() token.Token { return fp.Token }
func (fp *Float32Literal) String() string {
	return fp.Token.Literal
}

func (pe *PrefixExpression) expressionNode()         {}
func (pe *PrefixExpression) TokenLiteral() string    { return pe.Token.Literal }
func (pe *PrefixExpression) TokenValue() token.Token { return pe.Token }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) expressionNode()         {}
func (ie *InfixExpression) TokenLiteral() string    { return ie.Token.Literal }
func (ie *InfixExpression) TokenValue() token.Token { return ie.Token }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String() + " ")
	out.WriteString(ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (b *Boolean) expressionNode()         {}
func (b *Boolean) TokenLiteral() string    { return b.Token.Literal }
func (b *Boolean) TokenValue() token.Token { return b.Token }
func (b *Boolean) String() string          { return b.Token.Literal }

func (ie *IfExpression) expressionNode()         {}
func (ie *IfExpression) TokenLiteral() string    { return ie.Token.Literal }
func (ie *IfExpression) TokenValue() token.Token { return ie.Token }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()

}

func (bs *BlockStatement) statementNode()          {}
func (bs *BlockStatement) TokenLiteral() string    { return bs.Token.Literal }
func (bs *BlockStatement) TokenValue() token.Token { return bs.Token }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (fl *FunctionLiteral) expressionNode()         {}
func (fl *FunctionLiteral) TokenLiteral() string    { return fl.Token.Literal }
func (fl *FunctionLiteral) TokenValue() token.Token { return fl.Token }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for index, p := range fl.Parameters {
		params = append(params, p.String()+":"+fl.Type.ParameterTypes[index].String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")->")
	out.WriteString(fl.Type.ReturnType.String())
	out.WriteString("{")
	out.WriteString(fl.Body.String())
	out.WriteString("}")

	return out.String()
}

func (ce *CallExpression) expressionNode()         {}
func (ce *CallExpression) TokenLiteral() string    { return ce.Token.Literal }
func (ce *CallExpression) TokenValue() token.Token { return ce.Token }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

func (sl *StringLiteral) expressionNode()         {}
func (sl *StringLiteral) TokenLiteral() string    { return sl.Token.Literal }
func (sl *StringLiteral) TokenValue() token.Token { return sl.Token }
func (sl *StringLiteral) String() string          { return sl.Token.Literal }

func (al *ArrayLiteral) expressionNode()         {}
func (al *ArrayLiteral) TokenLiteral() string    { return al.Token.Literal }
func (al *ArrayLiteral) TokenValue() token.Token { return al.Token }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}

	for _, a := range al.Elements {
		elements = append(elements, a.String())
	}
	out.WriteString(al.Type.String())
	out.WriteString("{")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("}")

	return out.String()
}

func (ie *IndexExpression) expressionNode()         {}
func (ie *IndexExpression) TokenLiteral() string    { return ie.Token.Literal }
func (ie *IndexExpression) TokenValue() token.Token { return ie.Token }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("]")
	out.WriteString(")")

	return out.String()
}

func (ae *AccessExpression) expressionNode()         {}
func (ae *AccessExpression) TokenLiteral() string    { return ae.Token.Literal }
func (ae *AccessExpression) TokenValue() token.Token { return ae.Token }
func (ae *AccessExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Left.String())
	out.WriteString(".")
	out.WriteString(ae.Right.String())

	return out.String()
}

func (hl *HashLiteral) expressionNode()         {}
func (hl *HashLiteral) TokenLiteral() string    { return hl.Token.Literal }
func (hl *HashLiteral) TokenValue() token.Token { return hl.Token }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	keys := make([]Expression, 0, len(hl.Pairs))
	for k := range hl.Pairs {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return hl.Pairs[keys[i]].String() < hl.Pairs[keys[j]].String()
	})

	pairs := []string{}
	for _, key := range keys {
		pairs = append(pairs, key.String()+":"+hl.Pairs[key].String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (ds *DeclarationStatement) statementNode()          {}
func (ds *DeclarationStatement) TokenLiteral() string    { return ds.Token.Literal }
func (ds *DeclarationStatement) TokenValue() token.Token { return ds.Token }
func (ds *DeclarationStatement) String() string {
	var out bytes.Buffer

	if ds.IsConstant {
		out.WriteString("const ")
	}
	if ds.Type != nil {
		out.WriteString(ds.Type.String())
		out.WriteString(" ")
	}
	out.WriteString(ds.Name.String())
	out.WriteString(" = ")
	out.WriteString(ds.Value.String())
	out.WriteString(";")

	return out.String()
}

func (td *TypedDeclarationStatement) statementNode()          {}
func (td *TypedDeclarationStatement) TokenLiteral() string    { return td.Token.Literal }
func (td *TypedDeclarationStatement) TokenValue() token.Token { return td.Token }
func (td *TypedDeclarationStatement) String() string {
	var out bytes.Buffer

	out.WriteString(td.Type.String())
	out.WriteString(" ")
	out.WriteString(td.Name.String())
	out.WriteString(" = ")
	out.WriteString(td.Value.String())
	out.WriteString(";")

	return out.String()
}

func (ad *AssignmentDeclarationStatement) statementNode()          {}
func (ad *AssignmentDeclarationStatement) TokenLiteral() string    { return ad.Token.Literal }
func (ad *AssignmentDeclarationStatement) TokenValue() token.Token { return ad.Token }
func (ad *AssignmentDeclarationStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ad.Name.String())
	out.WriteString(" := ")
	out.WriteString(ad.Value.String())
	out.WriteString(";")

	return out.String()
}

func (vu *VariableUpdateStatement) statementNode()          {}
func (vu *VariableUpdateStatement) TokenLiteral() string    { return vu.Token.Literal }
func (vu *VariableUpdateStatement) TokenValue() token.Token { return vu.Token }
func (vu *VariableUpdateStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vu.Left.String())
	out.WriteString(" ")
	out.WriteString(vu.Operator)
	out.WriteString(" ")
	out.WriteString(vu.Right.String())

	return out.String()
}

func (ht HashType) typeNode()               {}
func (ht HashType) TokenLiteral() string    { return ht.Token.Literal }
func (ht HashType) TokenValue() token.Token { return ht.Token }
func (ht HashType) String() string {
	var out bytes.Buffer
	out.WriteString("map{ ")
	out.WriteString(ht.KeyType.String())
	out.WriteString(" -> ")
	out.WriteString(ht.ValueType.String())
	out.WriteString(" }")

	return out.String()
}

func (at ArrayType) typeNode()               {}
func (at ArrayType) TokenLiteral() string    { return at.Token.Literal }
func (at ArrayType) TokenValue() token.Token { return at.Token }

func (at ArrayType) String() string { return "[]" + at.ElementType.String() }

func (nt NamedType) typeNode()               {}
func (nt NamedType) TokenLiteral() string    { return nt.Token.Literal }
func (nt NamedType) TokenValue() token.Token { return nt.Token }
func (nt NamedType) String() string          { return nt.TypeName.String() }

func (ft FunctionType) typeNode()               {}
func (ft FunctionType) TokenLiteral() string    { return ft.Token.Literal }
func (ft FunctionType) TokenValue() token.Token { return ft.Token }
func (ft FunctionType) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range ft.ParameterTypes {
		params = append(params, p.String())
	}

	out.WriteString("function(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")->")
	out.WriteString(ft.ReturnType.String())

	return out.String()
}

func (fs *ForStatement) statementNode()          {}
func (fs *ForStatement) TokenLiteral() string    { return fs.Token.Literal }
func (fs *ForStatement) TokenValue() token.Token { return fs.Token }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString("for(")
	out.WriteString(fs.Initialization.String())
	out.WriteString(";")
	out.WriteString(fs.Condition.String())
	out.WriteString(";")
	out.WriteString(fs.AfterThought.String())
	out.WriteString("){")
	out.WriteString(fs.Body.String())
	out.WriteString("}")

	return out.String()
}

func (te *TernaryExpression) expressionNode()         {}
func (te *TernaryExpression) TokenLiteral() string    { return te.Token.Literal }
func (te *TernaryExpression) TokenValue() token.Token { return te.Token }
func (te *TernaryExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(te.Condition.String())
	out.WriteString(" ? ")
	out.WriteString(te.ValueIfTrue.String())
	out.WriteString(" : ")
	out.WriteString(te.ValueIfFalse.String())
	out.WriteString(")")

	return out.String()
}
