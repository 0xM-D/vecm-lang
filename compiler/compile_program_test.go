package compiler_test

import (
	"testing"
)

func TestCompileEmptyProgram(t *testing.T) {
	code := ``

	compileAndVerifyCode(code, t)
}
