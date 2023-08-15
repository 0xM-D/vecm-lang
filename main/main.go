package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/0xM-D/interpreter/evaluator"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/object"
	"github.com/0xM-D/interpreter/parser"
	"github.com/0xM-D/interpreter/repl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vecm",
	Short: "Vecm is a vector SIMD language interpreter",
	Long:  "Vecm is a programming language that emphasizes vector SIMD instructions and operations.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			filePath := args[0]
			runFile(filePath)
		} else {
			runRepl()
		}
	},
}

func runFile(filePath string) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(dat))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	println(evaluator.Eval(program, env).Inspect())
}

func runRepl() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s, this is the interpreter repl!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
