# structured-locale

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Coverage][coverage-svg]][coverage-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

A Go library for locale-related data and functionality with zero external dependencies.

## Features

- **BCP 47 Tag Parsing**: Parse and normalize language tags like `en-US`, `zh-Hans-CN`
- **Locale Fallback**: Automatic fallback chains (e.g., `fr-CA` → `fr` → `en`)
- **Translation Bundles**: Message translation with template variable substitution
- **CLDR Pluralization**: Full support for plural categories (zero, one, two, few, many, other)
- **Embedded Defaults**: Ships with translations for 6 locales (en, de, es, fr, ja, zh)
- **Override Support**: Customize any embedded data with your own translations

## Installation

```bash
go get github.com/grokify/structured-locale
```

## Usage

### Locale Tag Parsing

```go
import "github.com/grokify/structured-locale/locale"

// Parse a BCP 47 tag
tag, err := locale.Parse("zh-Hans-CN")
if err != nil {
    log.Fatal(err)
}

fmt.Println(tag.Language) // "zh"
fmt.Println(tag.Script)   // "Hans"
fmt.Println(tag.Region)   // "CN"
fmt.Println(tag.String()) // "zh-Hans-CN"

// Get parent for fallback
parent := tag.Parent() // "zh-Hans"
```

### Locale Fallback

```go
import "github.com/grokify/structured-locale/locale"

// Get fallback chain for a locale
chain := locale.FallbackChain("fr-CA", "en")
// Returns: ["fr-CA", "fr", "en"]
```

### Message Translation

```go
import "github.com/grokify/structured-locale/messages"

// Create a bundle with embedded defaults
bundle := messages.NewBundle()

// Get a localizer for French
loc := bundle.Localizer("fr")

// Simple translation
fmt.Println(loc.T("category.added"))     // "Ajouté"
fmt.Println(loc.T("category.fixed"))     // "Corrigé"

// Translation with variables
fmt.Println(loc.Tf("greeting", map[string]any{
    "Name": "Alice",
})) // "Bonjour, Alice!"

// Plural translation
fmt.Println(loc.Tn("items.count", 1))  // "1 élément"
fmt.Println(loc.Tn("items.count", 5))  // "5 éléments"
```

### Custom Translations

```go
import "github.com/grokify/structured-locale/messages"

// Load custom translations from a file
bundle := messages.NewBundle()
err := bundle.LoadFile("fr", "custom-fr.json")
if err != nil {
    log.Fatal(err)
}

// Or load from embedded data
//go:embed locales/*.json
var localesFS embed.FS

err = bundle.LoadFS(localesFS, "locales")
```

## Translation File Format

Translation files use a simple JSON format:

```json
[
  {
    "id": "category.added",
    "translation": "Added"
  },
  {
    "id": "items.count",
    "translation": {
      "one": "{{.Count}} item",
      "other": "{{.Count}} items"
    }
  }
]
```

See `schema/messages-v1.schema.json` for the full JSON Schema.

## Supported Locales

Built-in translations are provided for:

| Code | Language |
|------|----------|
| en | English |
| de | German |
| es | Spanish |
| fr | French |
| ja | Japanese |
| zh | Chinese |

## Packages

| Package | Description |
|---------|-------------|
| `locale` | BCP 47 tag parsing, normalization, fallback logic |
| `messages` | Translation bundles, pluralization, message formatting |

## Roadmap

See [PRD.md](PRD.md) for the full roadmap including planned packages for countries, phones, currencies, languages, dates, numbers, and addresses.

## License

MIT License - see [LICENSE](LICENSE) for details.

 [build-status-svg]: https://github.com/grokify/structured-locale/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/structured-locale/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/structured-locale/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/structured-locale/actions/workflows/lint.yaml
 [coverage-svg]: https://img.shields.io/badge/coverage-96.1%25-brightgreen
 [coverage-url]: https://github.com/grokify/structured-locale
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/structured-locale
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/structured-locale
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/structured-locale
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/structured-locale
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fstructured-locale
 [loc-svg]: https://tokei.rs/b1/github/grokify/structured-locale
 [repo-url]: https://github.com/grokify/structured-locale
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/structured-locale/blob/master/LICENSE
