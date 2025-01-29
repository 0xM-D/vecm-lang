package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/DustTheory/interpreter/compiler"
	"github.com/DustTheory/interpreter/module"
	"github.com/pkg/errors"
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
				compileFile(args[0])
			},
		}
	})
	return app.cmd
}

// func runFile(filePath string) {
// 	r, failedToLoad := runtime.NewRuntimeFromFile(filePath)
// 	if failedToLoad {
// 		return
// 	}

// 	runtimeError := r.Run()
// 	if runtimeError != nil {
// 		log.Println(runtimeError)
// 	}
// }

func compileFile(filePath string) {
	compiler, err := compiler.New()

	if err != nil {
		log.Fatal(error.Error(err))
		return
	}

	code, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	module, hasParserErrors := compiler.LoadModule(filePath, string(code))
	compiler.EntryModule = module

	if hasParserErrors {
		log.Fatal("Expected no parser errors, got some")
		return
	}

	// Create directory for temporary files
	tmpDir, err := os.MkdirTemp("", "vecm_output")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	finalExecutablePath, err := compileModule(compiler, module, tmpDir)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Copy the final executable to the current directory as "a.out"
	err = os.Rename(finalExecutablePath, "a.out")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

type ModuleType string

const (
	CoreModule   ModuleType = "core_module"
	LinkedModule ModuleType = "linked_module"
)

func compileLLIR(ir string, moduleType ModuleType, tmpDir string) (string, error) {
	pattern := getTempFilePattern(moduleType, "ll")

	sourceFile, err := writeToTemporaryFile(ir, pattern, tmpDir)
	if err != nil {
		return "", errors.Wrap(err, "Failed to write module to temporary file")
	}

	objFilePattern := getTempFilePattern(moduleType, "o")
	objFile, err := os.CreateTemp(tmpDir, objFilePattern)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create object file")
	}
	objFile.Close()

	//nolint:gosec // This is a toy project, security is not a concern
	cmd := exec.Command("llc", "-filetype=obj", "-o", objFile.Name(), sourceFile)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("error running llc for object file: %v\n%s", err, stderr.String()))
	}

	return objFile.Name(), nil
}

func compileModule(compiler *compiler.Compiler, mod *module.Module, tmpDir string) (string, error) {
	irModule, hasCompilerErrors := compiler.CompileModule(mod.ModuleKey)

	if hasCompilerErrors {
		compiler.PrintCompilerErrors()
		return "", errors.New("Expected no compiler errors, got some")
	}

	coreIRFilePath, coreFileCompileError := compileLLIR(irModule.CoreModule.String(), "core_module", tmpDir)
	if coreFileCompileError != nil {
		return "", errors.Wrap(coreFileCompileError, "Failed to write core module to temporary file")
	}

	linkModuleFiles := make([]string, 0, len(irModule.LinkedModules))
	for _, linkedModule := range irModule.LinkedModules {
		linkedIRFilePath, linkedFileCompileError := compileLLIR(linkedModule.String(), "linked_module", tmpDir)
		if linkedFileCompileError != nil {
			return "", errors.Wrap(linkedFileCompileError, "Failed to write linked module to temporary file")
		}
		linkModuleFiles = append(linkModuleFiles, linkedIRFilePath)
	}

	linkedModulePath, err := linkModules(coreIRFilePath, linkModuleFiles, tmpDir)

	if err != nil {
		return "", errors.Wrap(err, "Failed to link modules")
	}

	return linkedModulePath, nil
}

func writeToTemporaryFile(contents string, pattern string, tmpDir string) (string, error) {
	// Write core module to temporary file
	tempFile, err := os.CreateTemp(tmpDir, pattern)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create temporary file")
	}

	_, err = tempFile.WriteString(contents)

	if err != nil {
		return "", errors.Wrap(err, "Failed to write to temporary file")
	}
	tempFile.Close()

	return tempFile.Name(), nil
}

func linkModules(coreModulePath string, linkedModulePaths []string, tmpDir string) (string, error) {
	// Link core module with linked modules
	sourceFilesStr := fmt.Sprintf("%s ", coreModulePath)
	for _, linkedModulePath := range linkedModulePaths {
		sourceFilesStr += fmt.Sprintf("%s ", linkedModulePath)
	}

	//nolint:mnd // 8 is not a magic number
	linkedModulePath := fmt.Sprintf("%s/linked_module_%s.o", tmpDir, generateRandomString(8))

	// Split the sourceFilesStr into individual arguments
	sourceFilesArgs := strings.Fields(sourceFilesStr)

	// Create the command with the correct arguments
	for _, arg := range sourceFilesArgs {
		if strings.Contains(arg, ";") || strings.Contains(arg, "&") {
			return "", errors.New("invalid character in source file argument")
		}
	}
	//nolint:gosec // This is a command that is being run
	cmd := exec.Command("clang", append([]string{"-o", linkedModulePath}, sourceFilesArgs...)...)

	// Capture and print any errors from the command execution
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("error running ld for linked module: %v\n%s", err, string(output)))
	}

	return linkedModulePath, nil
}

func getTempFilePattern(moduleType ModuleType, ext string) string {
	switch moduleType {
	case CoreModule:
		return "core_module_*." + ext
	case LinkedModule:
		return "linked_module_*." + ext
	default:
		return ""
	}
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func main() {
	app := NewApplication()
	cmd := app.GetRootCmd()
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
