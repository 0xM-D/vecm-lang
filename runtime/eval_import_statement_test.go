package runtime

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
)

func TestImportStatement(t *testing.T) {
	testDir := t.TempDir()

	importedModuleFilePath := filepath.Join(testDir, filepath.Base("tempImportFile.vec"))
	importedModuleCode := `
		export integer1 := 1
		integer2 := 2
		export integer2

		export boolean1 := true

		export string1 := "chat is this real"
		export const constString = string1 + ", wowzers";

		const function1 = fn(x: int)->int { return x * 2; }
		export function1

		export fn main()->int {
			return 0;
		}
	`

	importFile, err := os.Create(importedModuleFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer importFile.Close()

	_, err = importFile.WriteString(importedModuleCode)
	if err != nil {
		t.Fatal(err)
	}

	err = importFile.Sync()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			fmt.Sprintf(`
				import integer1 from %q
				integer1
			`, importedModuleFilePath),
			big.NewInt(1),
		},
		{
			fmt.Sprintf(`
				import integer1, integer2 from %q
				integer1 + integer2
			`, importedModuleFilePath),
			big.NewInt(3),
		},
		{
			fmt.Sprintf(`
				import boolean1, string1 from %q
				boolean1
			`, importedModuleFilePath),
			true,
		},
		{
			fmt.Sprintf(`
				import constString from %q
				constString
			`, importedModuleFilePath),
			"chat is this real, wowzers",
		},
		{
			fmt.Sprintf(`
				import function1 from %q
				function1(13)
			`, importedModuleFilePath),
			big.NewInt(26),
		},
		{
			fmt.Sprintf(`
				import main from %q
				main()
			`, importedModuleFilePath),
			big.NewInt(0),
		},
	}

	for _, tt := range tests {
		result, err := testEval(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		testLiteralObject(t, result, tt.expected)
	}

}
