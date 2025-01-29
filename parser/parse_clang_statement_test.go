package parser_test

import (
	"strings"
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestCBlock(t *testing.T) {
	tests := []struct {
		input     string
		clangCode string
	}{
		{
			`CLang {
				#include <stdio.h>

				int printLine(char* str) {
					printf("%s\n", str);
				}
			}	
			
			printLine("Hello, World!");
			`,
			`
			#include <stdio.h>

				int printLine(char* str) {
					printf("%s\n", str);
				}
			`,
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) == 0 {
			t.Fatalf("program.Statement does not contain any statements")
		}

		clangStmt, isClangStmt := program.Statements[0].(*ast.CLangStatement)

		if !isClangStmt {
			t.Fatalf("program.Statements[0] is not a CLang statement. got=%T", program.Statements[0])
		}

		clangCode := strings.TrimSpace(clangStmt.CLangCode)
		testClangCode := strings.TrimSpace(tt.clangCode)

		if clangCode != testClangCode {
			t.Fatalf("CLang code is not %q. got=%q", testClangCode, clangCode)
		}
	}
}
