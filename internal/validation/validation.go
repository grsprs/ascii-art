package validation

import (
	"fmt"
)

const (
	// ASCII printable character range
	ASCIIPrintableStart = 32  // First printable ASCII character (space)
	ASCIIPrintableEnd   = 126 // Last printable ASCII character (~)
	
	// Valid banner names
	BannerStandard   = "standard"
	BannerShadow     = "shadow"
	BannerThinkertoy = "thinkertoy"
)

// validBanners is a pre-computed map for performance
var validBanners = map[string]bool{
	BannerStandard:   true,
	BannerShadow:     true,
	BannerThinkertoy: true,
}

// ValidateBanner checks if banner name is valid
func ValidateBanner(banner string) error {
	if !validBanners[banner] {
		return fmt.Errorf("invalid banner: %s (valid: standard, shadow, thinkertoy)", banner)
	}
	return nil
}

// ValidateInput checks if input contains only printable ASCII characters and newlines
func ValidateInput(input string) error {
	if len(input) == 0 {
		return nil // Empty input is valid
	}
	
	// Use byte-based iteration for accurate position reporting
	for i, char := range []byte(input) {
		if char != '\n' && (char < ASCIIPrintableStart || char > ASCIIPrintableEnd) {
			return fmt.Errorf("invalid character at position %d: %q (only ASCII %d-%d and newlines supported)", i, char, ASCIIPrintableStart, ASCIIPrintableEnd)
		}
	}
	return nil
}