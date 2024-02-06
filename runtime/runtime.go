package runtime

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/0xM-D/interpreter/module"
	"github.com/0xM-D/interpreter/object"
	"github.com/0xM-D/interpreter/parser"
)

type Runtime struct {
	Modules     map[string]*module.Module
	EntryModule *module.Module
}

func NewRuntimeFromFile(entryModulePath string) (*Runtime, bool) {
	runtime := &Runtime{Modules: map[string]*module.Module{}}
	entryModule, failedToLoad := runtime.loadModuleFromFile(entryModulePath)

	runtime.EntryModule = entryModule
	return runtime, failedToLoad
}

func NewRuntimeFromCode(code string) (*Runtime, bool) {
	runtime := &Runtime{Modules: map[string]*module.Module{}}
	entryModule, failedToLoad := runtime.loadModule("__entryPoint__", code)

	runtime.EntryModule = entryModule
	return runtime, failedToLoad
}

func (r *Runtime) Run() error {
	mainFunc := r.EntryModule.RootEnvironment.GetStore()["main"]

	if mainFunc == nil || !mainFunc.IsExported || !object.IsFunction(mainFunc.Object) {
		return fmt.Errorf("entry point main not exported from %s", r.EntryModule.ModuleKey)
	}

	_, runtimeError := r.ApplyFunction(mainFunc.Object, []object.Object{})
	if runtimeError != nil {
		return runtimeError
	}

	return nil
}

func (r *Runtime) loadModuleFromFile(modulePath string) (*module.Module, bool) {

	absolutePath, err := filepath.Abs(modulePath)
	if err != nil {
		fmt.Print(err.Error())
		return nil, true
	}

	cachedModule := r.Modules[absolutePath]
	if cachedModule != nil {
		return cachedModule, false
	}

	code, err := os.ReadFile(modulePath)
	if err != nil {
		fmt.Print(err.Error())
		return nil, true
	}

	return r.loadModule(absolutePath, string(code))
}

func (r *Runtime) loadModule(moduleKey string, code string) (*module.Module, bool) {

	module := module.ParseModule(moduleKey, code)
	if checkParserErrors(module.Parser) {
		return nil, true
	}

	r.Modules[moduleKey] = module

	r.Eval(module.Program, &module.RootEnvironment)

	return module, false
}

func checkParserErrors(p *parser.Parser) bool {
	errors := p.Errors()

	if len(errors) == 0 {
		return false
	}

	fmt.Printf("parser has %d errors\n", len(errors))
	for _, msg := range errors {
		fmt.Printf("parser error: %s", msg)
	}

	return true
}
