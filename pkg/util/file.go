package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrependCurrentDir(fileArg string) (string, error) {
	homeDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to resolve current directory: %w", err)
	}

	return filepath.Join(homeDir, fileArg), nil
}

func ReadFromPath(fileArg string) ([]byte, error) {
	content, err := os.ReadFile(fileArg)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return content, nil
}
