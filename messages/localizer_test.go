package messages

import (
	"testing"
)

func TestLocalizer_T(t *testing.T) {
	b := DefaultBundle()
	l := b.Localizer("en")

	got := l.T("changelog.title")
	if got != "Changelog" {
		t.Errorf("T('changelog.title') = %q, expected 'Changelog'", got)
	}

	// Missing key returns the key itself
	got = l.T("nonexistent.key")
	if got != "nonexistent.key" {
		t.Errorf("T('nonexistent.key') = %q, expected 'nonexistent.key'", got)
	}
}

func TestLocalizer_Tn(t *testing.T) {
	b := DefaultBundle()
	l := b.Localizer("en")

	// Test singular
	got := l.Tn("plural.releases", 1)
	if got != "1 release" {
		t.Errorf("Tn('plural.releases', 1) = %q, expected '1 release'", got)
	}

	// Test plural
	got = l.Tn("plural.releases", 5)
	if got != "5 releases" {
		t.Errorf("Tn('plural.releases', 5) = %q, expected '5 releases'", got)
	}

	// Test zero
	got = l.Tn("plural.releases", 0)
	if got != "0 releases" {
		t.Errorf("Tn('plural.releases', 0) = %q, expected '0 releases'", got)
	}
}

func TestLocalizer_Tf(t *testing.T) {
	b := DefaultBundle()
	l := b.Localizer("en")

	got := l.Tf("marker.versions_range", map[string]any{
		"From": "1.0.0",
		"To":   "1.5.0",
	})
	expected := "Versions 1.0.0 - 1.5.0"
	if got != expected {
		t.Errorf("Tf('marker.versions_range', ...) = %q, expected %q", got, expected)
	}
}

func TestLocalizer_Fallback(t *testing.T) {
	b := DefaultBundle()

	// Test fr-CA falls back to fr
	l := b.Localizer("fr-CA")
	got := l.T("changelog.title")
	if got != "Journal des modifications" {
		t.Errorf("fr-CA should fall back to fr, got %q", got)
	}
}

func TestLocalizer_French_Plural(t *testing.T) {
	b := DefaultBundle()
	l := b.Localizer("fr")

	// In French, 0 is singular (like 1)
	got := l.Tn("plural.releases", 0)
	// Should use singular form
	if got != "0 version" {
		t.Errorf("French Tn('plural.releases', 0) = %q, expected '0 version'", got)
	}

	got = l.Tn("plural.releases", 1)
	if got != "1 version" {
		t.Errorf("French Tn('plural.releases', 1) = %q, expected '1 version'", got)
	}

	got = l.Tn("plural.releases", 2)
	if got != "2 versions" {
		t.Errorf("French Tn('plural.releases', 2) = %q, expected '2 versions'", got)
	}
}

func TestLocalizer_Japanese_Plural(t *testing.T) {
	b := DefaultBundle()
	l := b.Localizer("ja")

	// Japanese has no plural distinctions
	got := l.Tn("plural.releases", 1)
	if got != "1件のリリース" {
		t.Errorf("Japanese Tn('plural.releases', 1) = %q, expected '1件のリリース'", got)
	}

	got = l.Tn("plural.releases", 5)
	if got != "5件のリリース" {
		t.Errorf("Japanese Tn('plural.releases', 5) = %q, expected '5件のリリース'", got)
	}
}
