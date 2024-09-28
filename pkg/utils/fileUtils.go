package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	Separator = string(os.PathSeparator)
)

// GetUserPath returns the current user's home directory
func GetUserPath() (string, error) {
	usrPath, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("failed to get user home dir: %v", usrPath)
	}
	return usrPath, err
}

// CreateDirectoriesIfNotExists create dirs by env path and file name
func CreateDirectoriesIfNotExists(envPath string) (*strings.Builder, error) {
	fullPath := strings.Builder{}
	paths := convertEnvPath(envPath)

	usrPath, err := GetUserPath()

	if err != nil {
		return &fullPath, err
	}

	fullPath.WriteString(usrPath)

	for _, path := range paths {
		fullPath.WriteString(Separator)
		fullPath.WriteString(path)
		err = createFileIfNotExists(fullPath.String())

		if err != nil {
			return &fullPath, fmt.Errorf("failed to create directory by path: %s", fullPath.String())
		}
	}

	return &fullPath, nil
}

// CloseFile closes file stream
func CloseFile(file *os.File) error {
	err := file.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %v", file)
	}

	return nil
}

// convertEnvPath converts a path from the environment into a suitable path for the OS.
// Env path separator must be "/"
func convertEnvPath(envPath string) []string {
	return strings.Split(envPath, "/")
}

func createFileIfNotExists(path string) error {
	if info, _ := os.Stat(path); info != nil {
		return nil
	}

	return os.Mkdir(path, 0755)
}
