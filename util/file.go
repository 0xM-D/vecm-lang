package util

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type File struct {
	Path  string
	Bytes []byte
}

func ReadFile(filePath string) (*File, error) {
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get absolute path")
	}

	code, err := os.ReadFile(absolutePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	return &File{absolutePath, code}, nil
}
