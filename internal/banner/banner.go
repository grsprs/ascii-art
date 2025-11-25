package banner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	asciiStart     = 32  // First printable ASCII character
	asciiEnd       = 126 // Last printable ASCII character
	charCount      = 95  // Total printable ASCII characters (126-32+1)
	linesPerChar   = 8   // Lines per character in banner
	separatorLines = 1   // Separator lines between characters
)

// Load loads a banner file and returns character mappings
func Load(filename string) (map[rune][]string, error) {
	if filename == "" {
		return nil, fmt.Errorf("empty filename")
	}
	// Validate file path - prevent all path traversal attacks
	cleanPath := filepath.FromSlash(filename)
	if !strings.HasSuffix(cleanPath, ".txt") {
		return nil, fmt.Errorf("invalid file extension")
	}
	// Block any path traversal attempts
	if filepath.IsAbs(cleanPath) || strings.Contains(cleanPath, "..") {
		return nil, fmt.Errorf("path traversal detected")
	}
	// Ensure path stays within expected directory structure
	if !strings.HasPrefix(cleanPath, "banners") {
		return nil, fmt.Errorf("path outside banners directory")
	}
	
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file")
	}

	// Handle Windows line endings efficiently
	replacer := strings.NewReplacer("\r\n", "\n", "\r", "\n")
	content := replacer.Replace(string(data))
	lines := strings.Split(content, "\n")
	
	// Skip first empty line if exists
	if len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}

	banner := make(map[rune][]string)

	// Validate total line count upfront
	expectedLines := charCount * (linesPerChar + separatorLines) - separatorLines
	if len(lines) < expectedLines {
		return nil, fmt.Errorf("invalid banner file format: expected at least %d lines, got %d", expectedLines, len(lines))
	}
	
	// Parse printable ASCII characters
	for i := 0; i < charCount; i++ {
		char := rune(asciiStart + i)
		start := i * (linesPerChar + separatorLines)
		if start+linesPerChar-1 >= len(lines) {
			return nil, fmt.Errorf("invalid banner file format: insufficient lines for character %c", char)
		}
		banner[char] = lines[start : start+linesPerChar]
	}

	return banner, nil
}