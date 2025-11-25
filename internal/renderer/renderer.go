package renderer

import (
	"strings"
)

const linesPerChar = 8 // Lines per character in ASCII art

// renderLine renders a single line of text to ASCII art
func renderLine(line string, banner map[rune][]string, result *strings.Builder) {
	for row := 0; row < linesPerChar; row++ {
		for _, char := range line {
			if art, exists := banner[char]; exists {
				if row < len(art) {
					result.WriteString(art[row])
				} else {
					// Fallback for insufficient art lines - use width of first line
					if len(art) > 0 {
						result.WriteString(strings.Repeat(" ", len(art[0])))
					} else {
						result.WriteString("      ") // Default fallback
					}
				}
			} else {
				// Use default spacing for missing characters
				result.WriteString("      ")
			}
		}
		result.WriteString("\n")
	}
}

// Render converts input text to ASCII art using the provided banner
func Render(input string, banner map[rune][]string) string {
	if len(banner) == 0 {
		return "" // Empty banner
	}
	
	// Replace \n with actual newlines
	input = strings.ReplaceAll(input, "\\n", "\n")
	
	var result strings.Builder
	// Pre-allocate based on ASCII art dimensions
	result.Grow(len(input) * 8 * 8) // 8 lines * ~8 chars width per character
	
	// Optimize for single-line input (common case)
	if !strings.Contains(input, "\n") {
		renderLine(input, banner, &result)
		return result.String()
	}
	
	// Multi-line processing
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			// Maintain consistent height for empty lines
			result.WriteString(strings.Repeat("\n", linesPerChar))
			continue
		}

		// Render this line
		renderLine(line, banner, &result)
	}

	return result.String()
}