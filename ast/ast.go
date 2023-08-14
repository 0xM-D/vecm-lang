package ast

import (
	"bytes"
	"strings"
)

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
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

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String() + " ")
	out.WriteString(oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
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

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
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

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
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

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}

	for _, a := range al.Elements {
		elements = append(elements, a.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
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

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (td *TypedDeclarationStatement) statementNode()       {}
func (td *TypedDeclarationStatement) TokenLiteral() string { return td.Token.Literal }
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

func (ad *AssignmentDeclarationStatement) statementNode()       {}
func (ad *AssignmentDeclarationStatement) TokenLiteral() string { return ad.Token.Literal }
func (ad *AssignmentDeclarationStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ad.Name.String())
	out.WriteString(" := ")
	out.WriteString(ad.Value.String())
	out.WriteString(";")

	return out.String()
}

func (vu *VariableUpdateStatement) statementNode()       {}
func (vu *VariableUpdateStatement) TokenLiteral() string { return vu.Token.Literal }
func (vu *VariableUpdateStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vu.Left.String())
	out.WriteString(" ")
	out.WriteString(vu.Operator)
	out.WriteString(" ")
	out.WriteString(vu.Right.String())

	return out.String()
}

func (ht HashType) typeNode()            {}
func (ht HashType) TokenLiteral() string { return ht.TokenLiteral() }
func (ht HashType) String() string {
	var out bytes.Buffer
	out.WriteString("map{ ")
	out.WriteString(ht.KeyType.String())
	out.WriteString(" -> ")
	out.WriteString(ht.ValueType.String())
	out.WriteString(" }")

	return out.String()
}

func (at ArrayType) typeNode()            {}
func (at ArrayType) TokenLiteral() string { return at.TokenLiteral() }
func (at ArrayType) String() string       { return at.ElementType.String() + "[]" }

func (it NamedType) typeNode()            {}
func (it NamedType) TokenLiteral() string { return it.TokenLiteral() }
func (it NamedType) String() string       { return it.TypeName.String() }

func (ft FunctionType) typeNode()            {}
func (ft FunctionType) TokenLiteral() string { return ft.TokenLiteral() }
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

func (rt ReturnType) String() string {
	if rt.Type == nil {
		return "void"
	}
	return (*rt.Type).String()
}
