package messages

import (
	"embed"
	"strings"
)

//go:embed locales/*.json
var defaultLocales embed.FS

// DefaultBundle returns a Bundle preloaded with embedded translations.
// Supports: en, fr, de, es, ja, zh.
func DefaultBundle() *Bundle {
	b := NewBundle("en")

	entries, err := defaultLocales.ReadDir("locales")
	if err != nil {
		return b
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}

		data, err := defaultLocales.ReadFile("locales/" + e.Name())
		if err != nil {
			continue
		}

		loc := strings.TrimSuffix(e.Name(), ".json")
		_ = b.AddLocale(loc, data)
	}

	return b
}
