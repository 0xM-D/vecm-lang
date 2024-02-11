package compiler

import "fmt"

func newError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
