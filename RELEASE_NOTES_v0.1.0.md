# structured-locale v0.1.0 Release Notes

**Release Date:** 2026-02-08

This is the initial release of `structured-locale`, a Go library designed for robust locale-related data and functionality without external dependencies. This version introduces core features for BCP 47 language tag parsing, intelligent locale fallback mechanisms, and a comprehensive message translation system with full CLDR pluralization support.

## Highlights

-   **BCP 47 Locale Parsing & Normalization**: Easily parse and normalize language tags like `en-US`, `zh-Hans-CN`.
-   **Automatic Locale Fallback**: Implement sophisticated fallback chains, e.g., `fr-CA` → `fr` → `en`.
-   **Message Translation Bundles**: Manage and translate messages with dynamic template variable substitution.
-   **CLDR Pluralization Support**: Full support for all plural categories (zero, one, two, few, many, other) ensures accurate translations for varying counts.
-   **Embedded Default Translations**: Comes with built-in translations for 6 key locales: English (en), German (de), Spanish (es), French (fr), Japanese (ja), and Chinese (zh).
-   **Extensible & Customizable**: Override embedded data or add your own translations with ease.

## New Features

-   **`locale` package**:
    -   BCP 47 tag parsing and normalization (`locale.Parse()`).
    -   Locale fallback chain generation (`locale.FallbackChain()`).
-   **`messages` package**:
    -   Translation bundle management (`messages.NewBundle()`, `messages.LoadFile()`, `messages.LoadFS()`).
    -   Localizer with `T()` for simple translations, `Tn()` for plural translations, and `Tf()` for translations with template variables.
    -   CLDR plural category support.
    -   Template variable substitution (`{{.Name}}` syntax).
    -   Embedded default translations for `en`, `de`, `es`, `fr`, `ja`, `zh`.
-   **JSON Schema**: Provided schema for message files (`schema/messages-v1.schema.json`) to ensure data consistency.

## Documentation

-   `README.md`: Comprehensive overview and usage examples for locale parsing and message translation.
-   `PRD.md`: Detailed roadmap outlining future packages and enhancements (countries, phones, currencies, etc.).

## Installation

```bash
go get github.com/grokify/structured-locale
```

## License

`structured-locale` is released under the MIT License.