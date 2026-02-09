package locale

// FallbackChain returns the fallback chain for a locale.
// Example: "fr-CA" with default "en" returns ["fr-CA", "fr", "en"].
// The chain always ends with the default locale if different from the input.
func FallbackChain(tag string, defaultLocale string) []string {
	var chain []string
	seen := make(map[string]bool)

	// Parse the input tag
	t, err := Parse(tag)
	if err != nil {
		// Invalid tag, just use default
		if defaultLocale != "" {
			return []string{defaultLocale}
		}
		return nil
	}

	// Walk up the parent chain
	for !t.IsZero() {
		s := t.String()
		if !seen[s] {
			chain = append(chain, s)
			seen[s] = true
		}
		t = t.Parent()
	}

	// Add default locale if not already in chain
	if defaultLocale != "" && !seen[defaultLocale] {
		chain = append(chain, defaultLocale)
	}

	return chain
}

// BestMatch finds the best matching locale from available locales.
// Returns the first match in the fallback chain, or the default locale if no match.
func BestMatch(requested string, available []string, defaultLocale string) string {
	if len(available) == 0 {
		return defaultLocale
	}

	// Build set of available locales (normalized)
	availableSet := make(map[string]string, len(available))
	for _, loc := range available {
		t, err := Parse(loc)
		if err != nil {
			continue
		}
		availableSet[t.String()] = loc // Map normalized to original
	}

	// Try each locale in the fallback chain
	chain := FallbackChain(requested, defaultLocale)
	for _, loc := range chain {
		if original, ok := availableSet[loc]; ok {
			return original
		}
	}

	// Fallback to default
	return defaultLocale
}
