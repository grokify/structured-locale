package messages

import (
	"github.com/grokify/structured-locale/locale"
)

// PluralCategory represents CLDR plural categories.
type PluralCategory string

const (
	PluralZero  PluralCategory = "zero"
	PluralOne   PluralCategory = "one"
	PluralTwo   PluralCategory = "two"
	PluralFew   PluralCategory = "few"
	PluralMany  PluralCategory = "many"
	PluralOther PluralCategory = "other"
)

// GetPluralCategory returns the CLDR plural category for a count in a locale.
// Implements common plural rules for major languages.
// For unsupported locales, falls back to English rules (one/other).
func GetPluralCategory(loc string, count int) PluralCategory {
	// Parse to get the base language
	t, err := locale.Parse(loc)
	if err != nil {
		return getPluralCategoryEnglish(count)
	}

	// Dispatch based on language
	switch t.Language {
	// Languages with only "other" (no singular form)
	case "ja", "ko", "zh", "vi", "th", "id", "ms":
		return PluralOther

	// Languages with one/other (Germanic, Romance, etc.)
	case "en", "de", "nl", "sv", "da", "no", "nb", "nn",
		"es", "pt", "it", "ca", "eu",
		"el", "he", "fi", "et", "hu", "tr":
		return getPluralCategoryEnglish(count)

	// French: 0 and 1 are singular
	case "fr":
		return getPluralCategoryFrench(count)

	// Slavic languages (Russian, Ukrainian, Polish, etc.)
	case "ru", "uk", "be":
		return getPluralCategorySlavic(count)

	// Polish has its own rules
	case "pl":
		return getPluralCategoryPolish(count)

	// Czech and Slovak
	case "cs", "sk":
		return getPluralCategoryCzech(count)

	// Arabic
	case "ar":
		return getPluralCategoryArabic(count)

	default:
		return getPluralCategoryEnglish(count)
	}
}

// getPluralCategoryEnglish returns plural category for English-like languages.
// one: n = 1
// other: everything else
func getPluralCategoryEnglish(n int) PluralCategory {
	if n == 1 {
		return PluralOne
	}
	return PluralOther
}

// getPluralCategoryFrench returns plural category for French.
// one: n = 0 or n = 1
// other: everything else
func getPluralCategoryFrench(n int) PluralCategory {
	if n == 0 || n == 1 {
		return PluralOne
	}
	return PluralOther
}

// getPluralCategorySlavic returns plural category for Russian/Ukrainian/Belarusian.
// one: n mod 10 = 1 and n mod 100 != 11
// few: n mod 10 in 2..4 and n mod 100 not in 12..14
// many: n mod 10 = 0 or n mod 10 in 5..9 or n mod 100 in 11..14
// other: (fractional numbers, not applicable for integers)
func getPluralCategorySlavic(n int) PluralCategory {
	mod10 := n % 10
	mod100 := n % 100

	if mod10 == 1 && mod100 != 11 {
		return PluralOne
	}
	if mod10 >= 2 && mod10 <= 4 && (mod100 < 12 || mod100 > 14) {
		return PluralFew
	}
	return PluralMany
}

// getPluralCategoryPolish returns plural category for Polish.
// one: n = 1
// few: n mod 10 in 2..4 and n mod 100 not in 12..14
// many: n != 1 and n mod 10 in 0..1 or n mod 10 in 5..9 or n mod 100 in 12..14
func getPluralCategoryPolish(n int) PluralCategory {
	if n == 1 {
		return PluralOne
	}

	mod10 := n % 10
	mod100 := n % 100

	if mod10 >= 2 && mod10 <= 4 && (mod100 < 12 || mod100 > 14) {
		return PluralFew
	}
	return PluralMany
}

// getPluralCategoryCzech returns plural category for Czech/Slovak.
// one: n = 1
// few: n in 2..4
// other: everything else
func getPluralCategoryCzech(n int) PluralCategory {
	if n == 1 {
		return PluralOne
	}
	if n >= 2 && n <= 4 {
		return PluralFew
	}
	return PluralOther
}

// getPluralCategoryArabic returns plural category for Arabic.
// zero: n = 0
// one: n = 1
// two: n = 2
// few: n mod 100 in 3..10
// many: n mod 100 in 11..99
// other: everything else
func getPluralCategoryArabic(n int) PluralCategory {
	if n == 0 {
		return PluralZero
	}
	if n == 1 {
		return PluralOne
	}
	if n == 2 {
		return PluralTwo
	}

	mod100 := n % 100
	if mod100 >= 3 && mod100 <= 10 {
		return PluralFew
	}
	if mod100 >= 11 {
		return PluralMany
	}
	return PluralOther
}
