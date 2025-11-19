package artefact

import (
	"fmt"
	"os"
)

// HandleDirectory reads the directory entries and returns their names.
func HandleDirectory(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("read directory %q: %w", path, err)
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	return names, nil
}

// HandleFile reads the entire file contents and returns them.
func HandleFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", path, err)
	}

	return data, nil
}

// HandleUnknown reports that the provided path type is not supported.
func HandleUnknown(path string) error {
	return fmt.Errorf("unsupported resource type for %s", path)
}
