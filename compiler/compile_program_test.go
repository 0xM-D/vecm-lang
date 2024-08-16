package compiler

import (
	"testing"
)

func TestCompileEmptyProgram(t *testing.T) {
	code := ``;

	compileAndVerifyCode(code, t)

} 