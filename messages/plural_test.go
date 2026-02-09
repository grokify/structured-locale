package messages

import (
	"testing"
)

func TestGetPluralCategory(t *testing.T) {
	tests := []struct {
		name     string
		locale   string
		count    int
		expected PluralCategory
	}{
		// English (one/other)
		{"en singular", "en", 1, PluralOne},
		{"en plural", "en", 0, PluralOther},
		{"en plural", "en", 2, PluralOther},
		{"en plural", "en", 100, PluralOther},

		// French (0 and 1 are singular)
		{"fr singular 0", "fr", 0, PluralOne},
		{"fr singular 1", "fr", 1, PluralOne},
		{"fr plural", "fr", 2, PluralOther},

		// Japanese (no plural)
		{"ja no plural", "ja", 1, PluralOther},
		{"ja no plural", "ja", 100, PluralOther},

		// Chinese (no plural)
		{"zh no plural", "zh", 1, PluralOther},
		{"zh-Hans no plural", "zh-Hans", 100, PluralOther},

		// Russian (complex rules)
		{"ru one", "ru", 1, PluralOne},
		{"ru one 21", "ru", 21, PluralOne},
		{"ru few 2", "ru", 2, PluralFew},
		{"ru few 3", "ru", 3, PluralFew},
		{"ru few 4", "ru", 4, PluralFew},
		{"ru many 0", "ru", 0, PluralMany},
		{"ru many 5", "ru", 5, PluralMany},
		{"ru many 11", "ru", 11, PluralMany},
		{"ru many 12", "ru", 12, PluralMany},
		{"ru many 100", "ru", 100, PluralMany},

		// Polish
		{"pl one", "pl", 1, PluralOne},
		{"pl few 2", "pl", 2, PluralFew},
		{"pl few 4", "pl", 4, PluralFew},
		{"pl many 0", "pl", 0, PluralMany},
		{"pl many 5", "pl", 5, PluralMany},
		{"pl many 12", "pl", 12, PluralMany},

		// Arabic
		{"ar zero", "ar", 0, PluralZero},
		{"ar one", "ar", 1, PluralOne},
		{"ar two", "ar", 2, PluralTwo},
		{"ar few 3", "ar", 3, PluralFew},
		{"ar few 10", "ar", 10, PluralFew},
		{"ar many 11", "ar", 11, PluralMany},
		{"ar many 99", "ar", 99, PluralMany},
		{"ar other 100", "ar", 100, PluralOther},

		// Region variants use base language
		{"en-US singular", "en-US", 1, PluralOne},
		{"fr-CA singular", "fr-CA", 0, PluralOne},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPluralCategory(tt.locale, tt.count)
			if got != tt.expected {
				t.Errorf("GetPluralCategory(%q, %d) = %q, expected %q",
					tt.locale, tt.count, got, tt.expected)
			}
		})
	}
}
