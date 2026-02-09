// Package locale provides BCP 47 language tag parsing and locale fallback logic.
package locale

import (
	"errors"
	"strings"
)

// Tag represents a parsed BCP 47 language tag.
type Tag struct {
	Language string // ISO 639-1 or 639-3 (required, lowercase)
	Script   string // ISO 15924 (optional, title case)
	Region   string // ISO 3166-1 alpha-2 (optional, uppercase)
}

// ErrInvalidTag is returned when a BCP 47 tag cannot be parsed.
var ErrInvalidTag = errors.New("invalid BCP 47 language tag")

// Parse parses a BCP 47 language tag string.
// Accepts formats like: "en", "en-US", "zh-Hans", "zh-Hans-CN".
// The input is case-insensitive; output is normalized.
func Parse(tag string) (Tag, error) {
	if tag == "" {
		return Tag{}, ErrInvalidTag
	}

	// Normalize separators: accept both - and _
	tag = strings.ReplaceAll(tag, "_", "-")
	parts := strings.Split(tag, "-")

	if len(parts) == 0 || len(parts[0]) < 2 {
		return Tag{}, ErrInvalidTag
	}

	t := Tag{
		Language: strings.ToLower(parts[0]),
	}

	// Validate language: 2-3 letters
	if len(t.Language) < 2 || len(t.Language) > 3 {
		return Tag{}, ErrInvalidTag
	}
	for _, c := range t.Language {
		if c < 'a' || c > 'z' {
			return Tag{}, ErrInvalidTag
		}
	}

	// Parse remaining parts
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		if len(part) == 0 {
			continue
		}

		// Script: 4 letters (e.g., "Hans", "Latn")
		if len(part) == 4 && t.Script == "" {
			t.Script = titleCase(part)
			continue
		}

		// Region: 2 letters (ISO 3166-1) or 3 digits (UN M.49)
		if len(part) == 2 && t.Region == "" {
			t.Region = strings.ToUpper(part)
			continue
		}
		if len(part) == 3 && t.Region == "" {
			// Check if it's all digits (UN M.49 code)
			allDigits := true
			for _, c := range part {
				if c < '0' || c > '9' {
					allDigits = false
					break
				}
			}
			if allDigits {
				t.Region = part
				continue
			}
		}

		// Ignore other subtags (variants, extensions, private use)
	}

	return t, nil
}

// String returns the normalized BCP 47 tag string.
func (t Tag) String() string {
	if t.Language == "" {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(t.Language)

	if t.Script != "" {
		sb.WriteString("-")
		sb.WriteString(t.Script)
	}

	if t.Region != "" {
		sb.WriteString("-")
		sb.WriteString(t.Region)
	}

	return sb.String()
}

// Parent returns the parent tag for fallback.
// For "fr-CA" returns "fr", for "zh-Hans-CN" returns "zh-Hans", for "en" returns empty Tag.
func (t Tag) Parent() Tag {
	if t.Region != "" {
		return Tag{Language: t.Language, Script: t.Script}
	}
	if t.Script != "" {
		return Tag{Language: t.Language}
	}
	return Tag{}
}

// IsZero returns true if the tag is empty/unset.
func (t Tag) IsZero() bool {
	return t.Language == ""
}

// MustParse parses a BCP 47 tag and panics on error.
// Use for known-good tags in initialization code.
func MustParse(tag string) Tag {
	t, err := Parse(tag)
	if err != nil {
		panic(err)
	}
	return t
}

// titleCase returns a string with the first letter uppercase and the rest lowercase.
// Used for script codes like "Hans", "Latn".
func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	lower := strings.ToLower(s)
	return strings.ToUpper(lower[:1]) + lower[1:]
}
