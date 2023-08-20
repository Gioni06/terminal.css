package build

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDirectoryExists(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Define a new directory path inside the temporary directory
	newDir := filepath.Join(tmpDir, "testdir")

	// 1. Test if the directory does not exist
	ensureDirectoryExists(newDir)
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		t.Errorf("Directory %s was not created", newDir)
	}

	// 2. Test if the directory already exists
	ensureDirectoryExists(newDir) // should not return any error or create a new directory
}
