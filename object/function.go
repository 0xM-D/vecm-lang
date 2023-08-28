package object

import (
	"bytes"
	"strings"

	"github.com/0xM-D/interpreter/ast"
)

type FunctionObjectType struct {
	ParameterTypes  []ObjectType
	ReturnValueType ObjectType
}

func (f FunctionObjectType) Signature() string {
	var out bytes.Buffer

	paramTypes := []string{}
	for _, p := range f.ParameterTypes {
		paramTypes = append(paramTypes, p.Signature())
	}

	out.WriteString("function(")
	out.WriteString(strings.Join(paramTypes, ", "))
	out.WriteString(") -> ")
	out.WriteString(f.ReturnValueType.Signature())

	return out.String()
}

func (f FunctionObjectType) Kind() ObjectKind              { return FunctionKind }
func (f FunctionObjectType) Builtins() *FunctionRepository { return FunctionKind.Builtins() }
func (f FunctionObjectType) IsConstant() bool              { return true }

type Function struct {
	FunctionObjectType
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return &f.FunctionObjectType }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()

}
