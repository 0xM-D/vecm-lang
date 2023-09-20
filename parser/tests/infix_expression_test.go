package parser_tests

import (
	"math/big"
	"testing"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/parser"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", big.NewInt(5), "+", big.NewInt(5)},
		{"5 - 5;", big.NewInt(5), "-", big.NewInt(5)},
		{"5 * 5;", big.NewInt(5), "*", big.NewInt(5)},
		{"5 / 5;", big.NewInt(5), "/", big.NewInt(5)},
		{"5 > 5;", big.NewInt(5), ">", big.NewInt(5)},
		{"5 < 5;", big.NewInt(5), "<", big.NewInt(5)},
		{"5 == 5;", big.NewInt(5), "==", big.NewInt(5)},
		{"5 != 5;", big.NewInt(5), "!=", big.NewInt(5)},
		{"foobar + barfoo;", TestIdentifier{"foobar"}, "+", TestIdentifier{"barfoo"}},
		{"foobar - barfoo;", TestIdentifier{"foobar"}, "-", TestIdentifier{"barfoo"}},
		{"foobar * barfoo;", TestIdentifier{"foobar"}, "*", TestIdentifier{"barfoo"}},
		{"foobar / barfoo;", TestIdentifier{"foobar"}, "/", TestIdentifier{"barfoo"}},
		{"foobar > barfoo;", TestIdentifier{"foobar"}, ">", TestIdentifier{"barfoo"}},
		{"foobar < barfoo;", TestIdentifier{"foobar"}, "<", TestIdentifier{"barfoo"}},
		{"foobar >= barfoo;", TestIdentifier{"foobar"}, ">=", TestIdentifier{"barfoo"}},
		{"foobar <= barfoo;", TestIdentifier{"foobar"}, "<=", TestIdentifier{"barfoo"}},
		{"foobar == barfoo;", TestIdentifier{"foobar"}, "==", TestIdentifier{"barfoo"}},
		{"foobar != barfoo;", TestIdentifier{"foobar"}, "!=", TestIdentifier{"barfoo"}},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
		{"true && false", true, "&&", false},
		{"false || false", false, "||", false},
		{"foobar & barfoo;", TestIdentifier{"foobar"}, "&", TestIdentifier{"barfoo"}},
		{"foobar | barfoo;", TestIdentifier{"foobar"}, "|", TestIdentifier{"barfoo"}},
		{"foobar ^ barfoo;", TestIdentifier{"foobar"}, "^", TestIdentifier{"barfoo"}},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not an ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		if !testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}
