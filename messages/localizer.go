package messages

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Localizer provides translation lookup for a specific locale.
type Localizer struct {
	bundle *Bundle
	locale string
}

// Locale returns the locale this localizer is configured for.
func (l *Localizer) Locale() string {
	return l.locale
}

// T translates a message ID to the localized string.
// Returns the ID itself if no translation is found.
func (l *Localizer) T(id string) string {
	m := l.bundle.GetMessage(l.locale, id)
	if m == nil {
		return id
	}
	return m.GetSingular()
}

// Tn translates a plural message ID with count.
// Returns the appropriate plural form based on the locale's plural rules.
func (l *Localizer) Tn(id string, count int) string {
	m := l.bundle.GetMessage(l.locale, id)
	if m == nil {
		return id
	}

	pt := m.GetPlural()
	if pt == nil {
		// Not a plural message, treat as singular
		return l.Tf(id, map[string]any{"Count": count})
	}

	// Get the appropriate plural category for this locale and count
	category := GetPluralCategory(l.locale, count)

	// Select the translation for this category
	var translation string
	switch category {
	case PluralZero:
		translation = pt.Zero
	case PluralOne:
		translation = pt.One
	case PluralTwo:
		translation = pt.Two
	case PluralFew:
		translation = pt.Few
	case PluralMany:
		translation = pt.Many
	default:
		translation = pt.Other
	}

	// Fall back to "other" if the category is empty
	if translation == "" {
		translation = pt.Other
	}

	// Substitute {{.Count}}
	return substituteVars(translation, map[string]any{"Count": count})
}

// Tf translates with template data.
// Variables in the format {{.Name}} are replaced with corresponding values.
func (l *Localizer) Tf(id string, data map[string]any) string {
	m := l.bundle.GetMessage(l.locale, id)
	if m == nil {
		return id
	}
	return substituteVars(m.GetSingular(), data)
}

// templateVarPattern matches {{.VarName}} patterns.
var templateVarPattern = regexp.MustCompile(`\{\{\s*\.(\w+)\s*\}\}`)

// substituteVars replaces {{.Name}} patterns with values from data.
func substituteVars(template string, data map[string]any) string {
	return templateVarPattern.ReplaceAllStringFunc(template, func(match string) string {
		// Extract variable name
		submatch := templateVarPattern.FindStringSubmatch(match)
		if len(submatch) < 2 {
			return match
		}
		varName := submatch[1]

		// Look up value (case-insensitive for flexibility)
		for k, v := range data {
			if strings.EqualFold(k, varName) {
				return formatValue(v)
			}
		}

		return match // Keep original if not found
	})
}

// formatValue converts a value to string for template substitution.
func formatValue(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}
