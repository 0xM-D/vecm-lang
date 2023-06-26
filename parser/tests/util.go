package parser_tests

import (
	"testing"

	"github.com/0xM-D/interpreter/parser"
)

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser errpr: %q", msg)
	}
	t.FailNow()
}
