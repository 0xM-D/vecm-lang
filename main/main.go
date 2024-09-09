package main

import (
	"log"
	"os"
	"sync"

	"github.com/DustTheory/interpreter/runtime"
	"github.com/spf13/cobra"
)

type Application struct {
	cmd  *cobra.Command
	once sync.Once
}

func NewApplication() *Application {
	return &Application{
		cmd:  nil,
		once: sync.Once{},
	}
}

func (app *Application) GetRootCmd() *cobra.Command {
	app.once.Do(func() {
		app.cmd = &cobra.Command{
			Use:   "vecm",
			Short: "Vecm is a vector SIMD language interpreter",
			Long:  "Vecm is a programming language that emphasizes vector SIMD instructions and operations.",
			Run: func(_ *cobra.Command, args []string) {
				runFile(args[0])
			},
		}
	})
	return app.cmd
}

func runFile(filePath string) {
	r, failedToLoad := runtime.NewRuntimeFromFile(filePath)
	if failedToLoad {
		return
	}

	runtimeError := r.Run()
	if runtimeError != nil {
		log.Println(runtimeError)
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
	app := NewApplication()
	cmd := app.GetRootCmd()
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
