package messages

import (
	"testing"
)

func TestBundle_AddLocale(t *testing.T) {
	b := NewBundle("en")

	jsonData := []byte(`{
		"messages": [
			{"id": "hello", "translation": "Hello"},
			{"id": "goodbye", "translation": "Goodbye"}
		]
	}`)

	err := b.AddLocale("en", jsonData)
	if err != nil {
		t.Fatalf("AddLocale failed: %v", err)
	}

	m := b.GetMessage("en", "hello")
	if m == nil {
		t.Fatal("GetMessage returned nil for 'hello'")
	}
	if m.GetSingular() != "Hello" {
		t.Errorf("Expected 'Hello', got %q", m.GetSingular())
	}
}

func TestBundle_Fallback(t *testing.T) {
	b := NewBundle("en")

	// Add English messages
	enData := []byte(`{
		"messages": [
			{"id": "hello", "translation": "Hello"},
			{"id": "only_en", "translation": "English only"}
		]
	}`)
	_ = b.AddLocale("en", enData)

	// Add French messages (missing "only_en")
	frData := []byte(`{
		"messages": [
			{"id": "hello", "translation": "Bonjour"}
		]
	}`)
	_ = b.AddLocale("fr", frData)

	// Test French gets French message
	m := b.GetMessage("fr", "hello")
	if m == nil || m.GetSingular() != "Bonjour" {
		t.Errorf("Expected 'Bonjour' for fr/hello, got %v", m)
	}

	// Test French falls back to English for missing message
	m = b.GetMessage("fr", "only_en")
	if m == nil || m.GetSingular() != "English only" {
		t.Errorf("Expected 'English only' for fr/only_en fallback, got %v", m)
	}

	// Test fr-CA falls back to fr, then to en
	m = b.GetMessage("fr-CA", "hello")
	if m == nil || m.GetSingular() != "Bonjour" {
		t.Errorf("Expected 'Bonjour' for fr-CA/hello fallback, got %v", m)
	}
}

func TestBundle_AddLocaleOverrides(t *testing.T) {
	b := NewBundle("en")

	// Add base messages
	baseData := []byte(`{
		"messages": [
			{"id": "hello", "translation": "Hello"},
			{"id": "goodbye", "translation": "Goodbye"}
		]
	}`)
	_ = b.AddLocale("en", baseData)

	// Add overrides
	overrideData := []byte(`{
		"messages": [
			{"id": "hello", "translation": "Hi there"}
		]
	}`)
	_ = b.AddLocaleOverrides("en", overrideData)

	// Check override
	m := b.GetMessage("en", "hello")
	if m == nil || m.GetSingular() != "Hi there" {
		t.Errorf("Expected 'Hi there' after override, got %v", m)
	}

	// Check non-overridden message still exists
	m = b.GetMessage("en", "goodbye")
	if m == nil || m.GetSingular() != "Goodbye" {
		t.Errorf("Expected 'Goodbye' to remain, got %v", m)
	}
}

func TestDefaultBundle(t *testing.T) {
	b := DefaultBundle()

	// Test that embedded locales are loaded
	locales := b.AvailableLocales()
	if len(locales) < 5 {
		t.Errorf("Expected at least 5 locales, got %d: %v", len(locales), locales)
	}

	// Test English
	m := b.GetMessage("en", "changelog.title")
	if m == nil || m.GetSingular() != "Changelog" {
		t.Errorf("Expected 'Changelog' for en/changelog.title, got %v", m)
	}

	// Test French
	m = b.GetMessage("fr", "changelog.title")
	if m == nil || m.GetSingular() != "Journal des modifications" {
		t.Errorf("Expected French title, got %v", m)
	}

	// Test Japanese
	m = b.GetMessage("ja", "changelog.title")
	if m == nil || m.GetSingular() != "変更履歴" {
		t.Errorf("Expected Japanese title, got %v", m)
	}
}
