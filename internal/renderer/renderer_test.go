package renderer

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	// Create a simple test banner
	testBanner := map[rune][]string{
		'A': {
			" _ ",
			"| |",
			"|_|",
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
		' ': {
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}

	tests := []struct {
		name     string
		input    string
		expected int // number of lines expected
	}{
		{"Single character", "A", 8},
		{"Empty string", "", 1},
		{"Space", " ", 8},
		{"Two characters", "AA", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Render(tt.input, testBanner)
			lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
			
			if tt.expected == 0 && result != "" {
				t.Errorf("Expected empty result for empty input, got %q", result)
			} else if tt.expected > 0 && len(lines) != tt.expected {
				t.Errorf("Expected %d lines, got %d", tt.expected, len(lines))
			}
		})
	}
}

func TestRenderMultiline(t *testing.T) {
	testBanner := map[rune][]string{
		'A': {"_", "|", " ", " ", " ", " ", " ", " "},
		'B': {"_", "|", " ", " ", " ", " ", " ", " "},
	}

	result := Render("A\\nB", testBanner)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	
	// Should have 8 lines for A + 1 empty line + 8 lines for B = 17 lines
	if len(lines) < 16 {
		t.Errorf("Expected at least 16 lines for multiline input, got %d", len(lines))
	}
}