package banner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	// Find project root by looking for go.mod
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	
	// Walk up directories to find project root
	projectRoot := wd
	for {
		if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(projectRoot)
		if parent == projectRoot {
			t.Skip("Could not find project root with go.mod")
		}
		projectRoot = parent
	}
	
	bannerPath := filepath.Join(projectRoot, "banners", "standard.txt")
	banner, err := Load(bannerPath)
	if err != nil {
		t.Skipf("Skipping test - banner file not found: %v", err)
	}

	// Test that we have the expected number of characters
	if len(banner) != charCount {
		t.Errorf("Expected %d characters, got %d", charCount, len(banner))
	}

	// Test specific characters
	if _, exists := banner['A']; !exists {
		t.Error("Character 'A' not found in banner")
	}

	if _, exists := banner[' ']; !exists {
		t.Error("Space character not found in banner")
	}

	// Test that each character has correct number of lines
	for char, lines := range banner {
		if len(lines) != linesPerChar {
			t.Errorf("Character %c has %d lines, expected %d", char, len(lines), linesPerChar)
		}
	}
}

func TestLoadInvalidFile(t *testing.T) {
	_, err := Load("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}