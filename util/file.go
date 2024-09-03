package util

import (
	"os"
	"path/filepath"
)

type File struct {
	Path  string
	Bytes []byte
}

func ReadFile(filePath string) (*File, error) {
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	code, err := os.ReadFile(absolutePath)
	if err != nil {
		return nil, err
	}

	return &File{absolutePath, code}, nil
}
