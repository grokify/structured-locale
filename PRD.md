# structured-locale Product Requirements Document

## Overview

structured-locale is a comprehensive Go library for locale-related data and functionality.
It provides structured JSON schemas and Go APIs for internationalization, localization,
and locale-specific data formatting.

## Packages

| Package | Description | Priority | Status |
|---------|-------------|----------|--------|
| `locale` | Core BCP 47 parsing, tag normalization, fallback logic | P0 | Implemented |
| `messages` | Translation bundles, pluralization, message formatting | P0 | Implemented |
| `countries` | ISO 3166 country codes, names, subdivisions, flags | P1 | Future |
| `phones` | E.164 formats, country calling codes, validation, formatting | P1 | Future |
| `currencies` | ISO 4217 codes, symbols, decimal places, formatting | P2 | Future |
| `languages` | ISO 639 language codes, native names, directionality | P2 | Future |
| `dates` | Date/time formats per locale, calendar systems | P2 | Future |
| `numbers` | Decimal separators, digit grouping, percent/currency formatting | P2 | Future |
| `addresses` | Address formats per country, field ordering | P3 | Future |

## Package Details

### locale (Core)

Core infrastructure shared across all packages.

- BCP 47 tag parsing and normalization (e.g., `en-US`, `fr-CA`, `zh-Hans-CN`)
- Locale fallback chains (e.g., `fr-CA` → `fr` → `en`)
- Tag matching and negotiation
- Region/script detection

### messages

Translation message bundles with pluralization support.

- go-i18n compatible JSON format
- CLDR plural categories (zero, one, two, few, many, other)
- Template variable substitution
- Embedded defaults with override support
- Fallback to parent locale or default

### countries

ISO 3166-1/2 country and subdivision data.

- Alpha-2, Alpha-3, and numeric codes
- Official and common names in multiple languages
- Subdivision codes (states, provinces, etc.)
- Emoji flags
- Continent/region grouping

### phones

E.164 phone number handling.

- Country calling codes (+1, +44, +81, etc.)
- National number formats per country
- Validation rules
- Formatting (national, international, E.164)
- Example numbers per type (mobile, landline, toll-free)

### currencies

ISO 4217 currency data.

- Currency codes (USD, EUR, JPY, etc.)
- Symbols ($, €, ¥, etc.)
- Decimal places (2 for USD, 0 for JPY, etc.)
- Name in multiple languages
- Formatting rules per locale

### languages

ISO 639 language data.

- ISO 639-1 (2-letter) and 639-3 (3-letter) codes
- Native names (endonyms)
- English names
- Script associations
- Text directionality (LTR/RTL)

### dates

Locale-specific date/time formatting.

- Date formats (short, medium, long, full)
- Time formats
- Date-time combinations
- First day of week
- Weekend days
- Month and day names

### numbers

Locale-specific number formatting.

- Decimal separators (. or ,)
- Digit grouping separators (, or . or space)
- Grouping patterns (3-digit, Indian numbering, etc.)
- Percent formatting
- Currency formatting integration

### addresses

Postal address formats per country.

- Field ordering (street, city, state, postal code)
- Required vs optional fields
- Label translations
- Postal code formats and validation

## Design Principles

1. **Structured Data**: All data defined in JSON schemas
2. **Embedded Defaults**: Ship with comprehensive default data
3. **Override Support**: Allow users to customize any data
4. **Lazy Loading**: Load only what's needed
5. **Zero Dependencies**: Core packages have no external dependencies
6. **Consistent API**: Same patterns across all packages

## Release Plan

- v0.1.0: `locale` + `messages` packages
- v0.2.0: `countries` package
- v0.3.0: `phones` package
- v0.4.0: `currencies` + `languages` packages
- v0.5.0: `dates` + `numbers` packages
- v0.6.0: `addresses` package
