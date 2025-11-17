package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	if input == "" {
		return
	}

	banner, err := loadBanner("banners/standard")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	output := renderText(input, banner)
	fmt.Print(output)
}

func loadBanner(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	banner := make([][]string, 95)
	
	for i := 0; i < 95; i++ {
		start := i * 9
		if start+7 < len(lines) {
			banner[i] = lines[start:start+8]
		}
	}
	
	return banner, nil
}

func renderText(text string, banner [][]string) string {
	// Replace \n with actual newlines
	text = strings.ReplaceAll(text, "\\n", "\n")
	
	if text == "\n" {
		return "\n"
	}

	lines := strings.Split(text, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range line {
				index := int(char) - 32
				if index >= 0 && index < len(banner) && banner[index] != nil && row < len(banner[index]) {
					result.WriteString(banner[index][row])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}