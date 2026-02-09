package locale

import (
	"reflect"
	"testing"
)

func TestFallbackChain(t *testing.T) {
	tests := []struct {
		name          string
		tag           string
		defaultLocale string
		expected      []string
	}{
		{
			name:          "simple language",
			tag:           "en",
			defaultLocale: "en",
			expected:      []string{"en"},
		},
		{
			name:          "language with region",
			tag:           "fr-CA",
			defaultLocale: "en",
			expected:      []string{"fr-CA", "fr", "en"},
		},
		{
			name:          "language with script and region",
			tag:           "zh-Hans-CN",
			defaultLocale: "en",
			expected:      []string{"zh-Hans-CN", "zh-Hans", "zh", "en"},
		},
		{
			name:          "same as default",
			tag:           "en",
			defaultLocale: "en",
			expected:      []string{"en"},
		},
		{
			name:          "invalid tag uses default",
			tag:           "invalid",
			defaultLocale: "en",
			expected:      []string{"en"},
		},
		{
			name:          "empty default",
			tag:           "fr-CA",
			defaultLocale: "",
			expected:      []string{"fr-CA", "fr"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FallbackChain(tt.tag, tt.defaultLocale)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FallbackChain(%q, %q) = %v, expected %v", tt.tag, tt.defaultLocale, got, tt.expected)
			}
		})
	}
}

func TestBestMatch(t *testing.T) {
	tests := []struct {
		name          string
		requested     string
		available     []string
		defaultLocale string
		expected      string
	}{
		{
			name:          "exact match",
			requested:     "fr-CA",
			available:     []string{"en", "fr", "fr-CA"},
			defaultLocale: "en",
			expected:      "fr-CA",
		},
		{
			name:          "fallback to parent",
			requested:     "fr-CA",
			available:     []string{"en", "fr"},
			defaultLocale: "en",
			expected:      "fr",
		},
		{
			name:          "fallback to default",
			requested:     "de-DE",
			available:     []string{"en", "fr"},
			defaultLocale: "en",
			expected:      "en",
		},
		{
			name:          "empty available returns default",
			requested:     "fr",
			available:     []string{},
			defaultLocale: "en",
			expected:      "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BestMatch(tt.requested, tt.available, tt.defaultLocale)
			if got != tt.expected {
				t.Errorf("BestMatch(%q, %v, %q) = %q, expected %q",
					tt.requested, tt.available, tt.defaultLocale, got, tt.expected)
			}
		})
	}
}
