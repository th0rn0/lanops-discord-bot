package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteDir_RemovesDirectory(t *testing.T) {
	dir := t.TempDir()
	// Create a file inside to confirm recursive delete
	if err := os.WriteFile(filepath.Join(dir, "test.txt"), []byte("data"), 0600); err != nil {
		t.Fatalf("setup: %v", err)
	}

	if err := DeleteDir(dir); err != nil {
		t.Fatalf("DeleteDir returned error: %v", err)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Error("directory still exists after DeleteDir")
	}
}

func TestDeleteDir_NonExistentPath(t *testing.T) {
	// os.RemoveAll is a no-op for non-existent paths — should not error
	if err := DeleteDir("/tmp/lanops-discord-bot-nonexistent-test-dir"); err != nil {
		t.Errorf("DeleteDir on non-existent path returned error: %v", err)
	}
}
