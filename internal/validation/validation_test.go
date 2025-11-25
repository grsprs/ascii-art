package validation

import (
	"testing"
)

func TestValidateBanner(t *testing.T) {
	tests := []struct {
		name      string
		banner    string
		shouldErr bool
	}{
		{"standard_banner", "standard", false},
		{"shadow_banner", "shadow", false},
		{"thinkertoy_banner", "thinkertoy", false},
		{"invalid_banner", "invalid", true},
		{"empty_banner", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBanner(tt.banner)
			if tt.shouldErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		shouldErr bool
	}{
		{"Valid ASCII", "Hello World!", false},
		{"Valid with newline", "Hello\nWorld", false},
		{"Valid numbers", "123", false},
		{"Valid symbols", "!@#$%^&*()", false},
		{"Invalid unicode", "Hello 世界", true},
		{"Invalid control char", "Hello\t", true},
		{"Invalid null char", "Hello\x00", true},
		{"Invalid DEL char", "Hello\x7F", true},
		{"Invalid char below 32", "Hello\x1F", true},
		{"Invalid char above 126", "Hello\x80", true},
		{"Empty input", "", false},
		{"Only newlines", "\n\n\n", false},
		{"Mixed valid chars", "Hello123!@#$%^&*()_+-=[]{}|;':,.<>?", false},
		{"Boundary chars", " ~", false}, // ASCII 32 and 126
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.input)
			if tt.shouldErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}