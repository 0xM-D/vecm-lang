package external_test

import (
	"testing"

	"github.com/DustTheory/interpreter/external"
	"github.com/DustTheory/interpreter/token"
)

// func TestImportClibraryFromPath(t *testing.T) {
// 	code := `
// 		int double_double(double d);
// 	`

// 	file, err := os.CreateTemp("", "*.c")
// 	if err != nil {
// 		t.Fatalf("Failed to create temporary file")
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(code)
// 	if err != nil {
// 		t.Fatalf("Failed to write code to file")
// 	}

// 	paths := []string{file.Name()}
// 	token := token.Token{
// 		Type:    token.Import,
// 		Literal: "import",
// 		Linen:   0,
// 		Coln:    0,
// 	}
// 	funcs, err := external.GetCLibraryFunctionsFromPath(paths, token)
// 	if err != nil {
// 		t.Fatalf("Failed to import C library: %s", err)
// 	}

// 	for _, f := range funcs {
// 		t.Logf("Function: %s", f.String())
// 	}
// }

func TestImportCStdlib(t *testing.T) {
	paths := []string{"math.h"}
	token := token.Token{
		Type:    token.Import,
		Literal: "import",
		Linen:   0,
		Coln:    0,
	}
	funcs, err := external.ImportCStdlib(paths, token)
	if err != nil {
		t.Fatalf("Failed to import C library: %s", err)
	}

	println("Funcs: %v", funcs)

	for _, f := range funcs {
		println("Function: %s", f.String())
	}

	println("Functions imported successfully")
}
