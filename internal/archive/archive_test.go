package archive

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateExtractArchive(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "archive-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	file1Path := filepath.Join(tempDir, "file1.txt")
	file2Path := filepath.Join(tempDir, "file2.txt")

	if err := os.WriteFile(file1Path, []byte("file1 content"), 0644); err != nil {
		t.Fatalf("failed to create file1: %v", err)
	}
	if err := os.WriteFile(file2Path, []byte("file2 content"), 0644); err != nil {
		t.Fatalf("failed to create file2: %v", err)
	}

	archiveData, err := CreateArchive([]string{file1Path, file2Path})
	if err != nil {
		t.Fatalf("failed to create archive: %v", err)
	}

	extractDir := filepath.Join(tempDir, "extracted")
	if err := os.Mkdir(extractDir, 0755); err != nil {
		t.Fatalf("failed to create extract dir: %v", err)
	}

	if err := ExtractArchive(archiveData, extractDir); err != nil {
		t.Fatalf("failed to extract archive: %v", err)
	}

	// Verify file1
	f1, err := os.ReadFile(filepath.Join(extractDir, "file1.txt"))
	if err != nil {
		t.Fatalf("failed to read extracted file1: %v", err)
	}
	if !bytes.Equal(f1, []byte("file1 content")) {
		t.Fatalf("file1 content mismatch")
	}

	// Verify file2
	f2, err := os.ReadFile(filepath.Join(extractDir, "file2.txt"))
	if err != nil {
		t.Fatalf("failed to read extracted file2: %v", err)
	}
	if !bytes.Equal(f2, []byte("file2 content")) {
		t.Fatalf("file2 content mismatch")
	}
}
