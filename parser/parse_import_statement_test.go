package parser

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
)

func TestImportStatement(t *testing.T) {

	tests := []struct {
		input                     string
		expectedImportPath        string
		expectedImportAll         bool
		expectedImportIdentifiers []string
	}{
		{`import a, b, c from "./file.vec"`, "./file.vec", false, []string{"a", "b", "c"}},
		{`import * from "../../file.vec"`, "../../file.vec", true, []string{}},
		{`import math from "./math.vec"`, "./math.vec", false, []string{"math"}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		importStmt, ok := program.Statements[0].(*ast.ImportStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not an import statement. got=%T", program.Statements[0])
		}

		if importStmt.ImportPath != tt.expectedImportPath {
			t.Fatalf("import path is not %q. got=%q", tt.expectedImportPath, importStmt.ImportPath)
		}

		if importStmt.ImportAll != tt.expectedImportAll {
			if tt.expectedImportAll {
				t.Fatalf("Expected import statement to import *")
			} else {
				t.Fatalf("Expected import statement to not import *")
			}
		}

		if len(importStmt.ImportedIdentifiers) != len(tt.expectedImportIdentifiers) {
			t.Fatalf("Expected import statement to import %d identifiers. got=%d", len(tt.expectedImportIdentifiers), len(importStmt.ImportedIdentifiers))
		}

		for i, identifier := range importStmt.ImportedIdentifiers {
			testIdentifier(t, identifier, tt.expectedImportIdentifiers[i])
		}

	}
}
