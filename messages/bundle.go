package messages

import (
	"github.com/grokify/structured-locale/locale"
)

// Bundle holds messages for multiple locales with fallback support.
type Bundle struct {
	defaultLocale string
	locales       map[string]*MessageSet
}

// NewBundle creates a bundle with the specified default locale.
func NewBundle(defaultLocale string) *Bundle {
	return &Bundle{
		defaultLocale: defaultLocale,
		locales:       make(map[string]*MessageSet),
	}
}

// DefaultLocale returns the bundle's default locale.
func (b *Bundle) DefaultLocale() string {
	return b.defaultLocale
}

// AddLocale adds messages for a locale from JSON data.
// Replaces any existing messages for this locale.
func (b *Bundle) AddLocale(loc string, data []byte) error {
	mf, err := ParseMessagesJSON(data)
	if err != nil {
		return err
	}

	// Normalize the locale tag
	t, err := locale.Parse(loc)
	if err != nil {
		return err
	}
	normalized := t.String()

	ms := NewMessageSet(normalized)
	for i := range mf.Messages {
		ms.Set(&mf.Messages[i])
	}

	b.locales[normalized] = ms
	return nil
}

// AddLocaleOverrides merges override messages into an existing locale.
// Only adds/updates the specified messages; others are unchanged.
func (b *Bundle) AddLocaleOverrides(loc string, data []byte) error {
	mf, err := ParseMessagesJSON(data)
	if err != nil {
		return err
	}

	// Normalize the locale tag
	t, err := locale.Parse(loc)
	if err != nil {
		return err
	}
	normalized := t.String()

	ms, ok := b.locales[normalized]
	if !ok {
		ms = NewMessageSet(normalized)
		b.locales[normalized] = ms
	}

	for i := range mf.Messages {
		ms.Set(&mf.Messages[i])
	}

	return nil
}

// Localizer returns a Localizer for the specified locale with fallback.
func (b *Bundle) Localizer(loc string) *Localizer {
	return &Localizer{
		bundle: b,
		locale: loc,
	}
}

// GetMessage retrieves a message by ID for the given locale.
// Uses fallback chain if not found in the requested locale.
func (b *Bundle) GetMessage(loc string, id string) *Message {
	chain := locale.FallbackChain(loc, b.defaultLocale)

	for _, l := range chain {
		if ms, ok := b.locales[l]; ok {
			if m := ms.Get(id); m != nil {
				return m
			}
		}
	}

	return nil
}

// AvailableLocales returns the list of loaded locales.
func (b *Bundle) AvailableLocales() []string {
	result := make([]string, 0, len(b.locales))
	for loc := range b.locales {
		result = append(result, loc)
	}
	return result
}
