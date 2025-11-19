package artefact

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHandleDirectory(t *testing.T) {
	dir := t.TempDir()

	files := []string{"alpha.txt", "beta.txt"}
	for _, name := range files {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(name), 0o600); err != nil {
			t.Fatalf("write temp file: %v", err)
		}
	}

	entries, err := HandleDirectory(dir)
	if err != nil {
		t.Fatalf("HandleDirectory returned error: %v", err)
	}

	if len(entries) != len(files) {
		t.Fatalf("HandleDirectory returned %d entries, want %d", len(entries), len(files))
	}

	for _, want := range files {
		found := false
		for _, got := range entries {
			if got == want {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("HandleDirectory result missing %s", want)
		}
	}
}

func TestHandleFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file.txt")
	content := []byte("hello world")
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	data, err := HandleFile(path)
	if err != nil {
		t.Fatalf("HandleFile returned error: %v", err)
	}

	if string(data) != string(content) {
		t.Fatalf("HandleFile returned %q, want %q", string(data), string(content))
	}
}

func TestHandleUnknown(t *testing.T) {
	if err := HandleUnknown("/path/to/socket"); err == nil {
		t.Fatal("HandleUnknown expected to return an error")
	}
}
