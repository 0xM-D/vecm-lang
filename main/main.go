package main

import (
	"fmt"
	"os"

	"github.com/DustTheory/interpreter/runtime"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vecm",
	Short: "Vecm is a vector SIMD language interpreter",
	Long:  "Vecm is a programming language that emphasizes vector SIMD instructions and operations.",
	Run: func(cmd *cobra.Command, args []string) {
		runFile(args[0])
	},
}

func runFile(filePath string) {

	r, failedToLoad := runtime.NewRuntimeFromFile(filePath)
	if failedToLoad {
		return
	}

	runtimeError := r.Run()
	if runtimeError != nil {
		fmt.Println(runtimeError)
	}
}

// func compileFile(filePath string) {

// 	compiler, error := compiler.InitializeCompiler()
// 	if error != nil {
// 		fmt.Errorf(error.Error())
// 		return
// 	}

// 	compiler.LoadEntryModuleFromFile(filePath)
// }

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
