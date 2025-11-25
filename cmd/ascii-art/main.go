package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/grsprs/ascii-art/internal/banner"
	"github.com/grsprs/ascii-art/internal/renderer"
	"github.com/grsprs/ascii-art/internal/validation"
)

const maxInputSize = 1024 * 1024 // 1MB limit

// validateAbsolutePath validates that a path stays within the working directory
func validateAbsolutePath(cleanPath string) error {
	if cleanPath == "" {
		return fmt.Errorf("empty path")
	}
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return fmt.Errorf("invalid file path")
	}
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get working directory")
	}
	if !strings.HasPrefix(absPath, cwd) {
		return fmt.Errorf("path outside working directory")
	}
	return nil
}

// validateFilePath checks if a file path is safe and within allowed directories
func validateFilePath(path string) error {
	if path == "" {
		return fmt.Errorf("empty path")
	}
	cleanPath := filepath.FromSlash(path)
	if filepath.IsAbs(cleanPath) || strings.Contains(cleanPath, "..") || strings.HasPrefix(cleanPath, "/") {
		return fmt.Errorf("invalid file path")
	}
	// Use filepath.Rel for robust path validation
	rel, err := filepath.Rel(".", cleanPath)
	if err != nil || strings.HasPrefix(rel, "..") {
		return fmt.Errorf("file path outside allowed directory")
	}
	return nil
}

func main() {
	// Define flags
	bannerFlag := flag.String("banner", "standard", "Banner style: standard, shadow, thinkertoy")
	fileFlag := flag.String("file", "", "Read input from file")
	outputFlag := flag.String("output", "", "Write output to file")
	versionFlag := flag.Bool("version", false, "Show version information")
	helpFlag := flag.Bool("help", false, "Show help")
	flag.Parse()
	
	// Handle version flag
	if *versionFlag {
		fmt.Println("ascii-art version 1.0.0")
		return
	}
	
	// Handle help flag
	if *helpFlag {
		flag.Usage()
		return
	}

	// Get input
	var input string
	if fileFlag != nil && *fileFlag != "" {
		// Validate file path
		if err := validateFilePath(*fileFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(2)
		}
		cleanPath := filepath.FromSlash(*fileFlag)
		
		// Additional security check
		if err := validateAbsolutePath(cleanPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(2)
		}
		
		// Check file size before reading
		info, err := os.Stat(cleanPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file\n")
			os.Exit(2)
		}

		if info.Size() > maxInputSize {
			fmt.Fprintf(os.Stderr, "Error: file too large (max 1MB)\n")
			os.Exit(2)
		}
		
		// Read from file
		data, err := os.ReadFile(cleanPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file\n")
			os.Exit(2)
		}
		input = strings.TrimRight(string(data), "\r\n")
	} else if len(flag.Args()) == 1 {
		// Read from command line argument
		input = flag.Args()[0]
		// Check input size
		if len(input) > maxInputSize {
			fmt.Fprintf(os.Stderr, "Error: input too large (max 1MB)\n")
			os.Exit(2)
		}
	} else if len(flag.Args()) > 1 {
		// Too many arguments
		fmt.Fprintf(os.Stderr, "Error: too many arguments provided\n")
		fmt.Fprintf(os.Stderr, "Usage: ascii-art [flags] <text> OR ascii-art -file <filename>\n")
		os.Exit(1)
	} else {
		// No arguments provided
		fmt.Fprintf(os.Stderr, "Error: no input provided\n")
		fmt.Fprintf(os.Stderr, "Usage: ascii-art [flags] <text> OR ascii-art -file <filename>\n")
		fmt.Fprintf(os.Stderr, "Use -help for more information\n")
		os.Exit(1)
	}

	// Handle empty string
	if input == "" {
		return
	}

	// Handle newline only - check before processing
	if input == "\\n" {
		fmt.Print("\n")
		return
	}

	// Validate banner
	if err := validation.ValidateBanner(*bannerFlag); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Validate input
	if err := validation.ValidateInput(input); err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid input\n")
		os.Exit(1)
	}

	// Load banner with secure path
	bannerFile := filepath.Join("banners", *bannerFlag+".txt")
	bannerData, err := banner.Load(bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner\n")
		os.Exit(3)
	}
	if bannerData == nil {
		fmt.Fprintf(os.Stderr, "Error: banner data is nil\n")
		os.Exit(3)
	}

	// Render text
	output := renderer.Render(input, bannerData)
	
	// Write output
	if outputFlag != nil && *outputFlag != "" {
		// Validate output path using existing function
		if err := validateFilePath(*outputFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(2)
		}
		cleanOutput := filepath.FromSlash(*outputFlag)
		// Additional security check
		if err := validateAbsolutePath(cleanOutput); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(2)
		}
		
		// Write to file
		if err := os.WriteFile(cleanOutput, []byte(output), 0o600); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file\n")
			os.Exit(2)
		}
	} else {
		// Write to stdout
		fmt.Print(output)
	}
}