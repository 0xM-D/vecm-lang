package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/0xM-D/interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s, this is the interpreter repl!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
