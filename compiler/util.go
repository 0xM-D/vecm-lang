package compiler

import (
	"github.com/llir/llvm/ir/constant"
)

// func getConcatenatedParserErrors(errors []string) error {
// 	var errs bytes.Buffer

// 	errs.WriteString(fmt.Sprintf("parser has %d errors\n", len(errors)))
// 	for _, msg := range errors {
// 		errs.WriteString(fmt.Sprintf("parser error: %s", msg))
// 	}

// 	return fmt.Fatalf(errs.String())
// }

func nativeBoolToLLVMBool(b bool) *constant.Int {
	if b {
		return constant.True
	} else {
		return constant.False
	}
}
