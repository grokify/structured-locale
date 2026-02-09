package locale

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected Tag
		wantErr  bool
	}{
		// Basic language codes
		{"en", Tag{Language: "en"}, false},
		{"fr", Tag{Language: "fr"}, false},
		{"zh", Tag{Language: "zh"}, false},

		// Language with region
		{"en-US", Tag{Language: "en", Region: "US"}, false},
		{"fr-CA", Tag{Language: "fr", Region: "CA"}, false},
		{"pt-BR", Tag{Language: "pt", Region: "BR"}, false},

		// Language with script
		{"zh-Hans", Tag{Language: "zh", Script: "Hans"}, false},
		{"zh-Hant", Tag{Language: "zh", Script: "Hant"}, false},

		// Language with script and region
		{"zh-Hans-CN", Tag{Language: "zh", Script: "Hans", Region: "CN"}, false},
		{"zh-Hant-TW", Tag{Language: "zh", Script: "Hant", Region: "TW"}, false},

		// Case normalization
		{"EN", Tag{Language: "en"}, false},
		{"EN-us", Tag{Language: "en", Region: "US"}, false},
		{"zh-HANS-cn", Tag{Language: "zh", Script: "Hans", Region: "CN"}, false},

		// Underscore separator (common alternative)
		{"en_US", Tag{Language: "en", Region: "US"}, false},
		{"zh_Hans_CN", Tag{Language: "zh", Script: "Hans", Region: "CN"}, false},

		// Three-letter language codes
		{"jpn", Tag{Language: "jpn"}, false},

		// Invalid inputs
		{"", Tag{}, true},
		{"x", Tag{}, true},
		{"123", Tag{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.expected {
				t.Errorf("Parse(%q) = %+v, expected %+v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestTag_String(t *testing.T) {
	tests := []struct {
		tag      Tag
		expected string
	}{
		{Tag{Language: "en"}, "en"},
		{Tag{Language: "en", Region: "US"}, "en-US"},
		{Tag{Language: "zh", Script: "Hans"}, "zh-Hans"},
		{Tag{Language: "zh", Script: "Hans", Region: "CN"}, "zh-Hans-CN"},
		{Tag{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			got := tt.tag.String()
			if got != tt.expected {
				t.Errorf("Tag%+v.String() = %q, expected %q", tt.tag, got, tt.expected)
			}
		})
	}
}

func TestTag_Parent(t *testing.T) {
	tests := []struct {
		tag      Tag
		expected Tag
	}{
		{Tag{Language: "en", Region: "US"}, Tag{Language: "en"}},
		{Tag{Language: "zh", Script: "Hans", Region: "CN"}, Tag{Language: "zh", Script: "Hans"}},
		{Tag{Language: "zh", Script: "Hans"}, Tag{Language: "zh"}},
		{Tag{Language: "en"}, Tag{}},
		{Tag{}, Tag{}},
	}

	for _, tt := range tests {
		t.Run(tt.tag.String(), func(t *testing.T) {
			got := tt.tag.Parent()
			if got != tt.expected {
				t.Errorf("Tag%+v.Parent() = %+v, expected %+v", tt.tag, got, tt.expected)
			}
		})
	}
}
