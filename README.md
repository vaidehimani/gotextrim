# 📝 gotextrim

A lightweight Go utility library for intelligently trimming text strings. This library provides an easy way to trim text to specified lengths while either preserving whole words or cutting mid-word as needed. It also handles punctuation appropriately. Perfect for summaries, previews, and UI components where clean text presentation matters.

[![Go Report Card](https://goreportcard.com/badge/github.com/vaidehimani/gotextrim)](https://goreportcard.com/report/github.com/vaidehimani/gotextrim)
[![GoDoc](https://godoc.org/github.com/vaidehimani/gotextrim?status.svg)](https://godoc.org/github.com/vaidehimani/gotextrim)
[![license](https://img.shields.io/github/license/vaidehimani/gotextrim)](LICENSE)
[![Build Status](https://github.com/vaidehimani/gotextrim/workflows/CI/badge.svg)](https://github.com/vaidehimani/gotextrim/actions)

## 💡 Examples

```go
// Returns original string if shorter than maxLength
fmt.Println(trimmer.SmartTrim("Short text", 20, nil))
// Output: "Short text"

// Custom suffix
options := &trimmer.SmartTrimOptions{
    Suffix: " [more]",
}
fmt.Println(trimmer.SmartTrim("This is a long sentence", 15, options))
// Output: "This is [more]"

// Cut words in the middle (disable whole word preservation)
fmt.Println(trimmer.SmartTrim("This is a long sentence", 15, nil, 
    trimmer.WithPreserveWholeWords(false)))
// Output: "This is a lo..."

// Remove trailing punctuation
fmt.Println(trimmer.SmartTrim("This is a sentence, with punctuation.", 22, nil,
    trimmer.WithPreservePunctuation(false)))
// Output: "This is a sentence..."

// Using functional options pattern (recommended)
fmt.Println(trimmer.SmartTrim("This is a long sentence", 18, nil,
    trimmer.WithSuffix(" [read more]"),
    trimmer.WithPreserveWholeWords(true)))
// Output: "This [read more]"

// Set empty suffix with functional option
fmt.Println(trimmer.SmartTrim("Short text", 6, nil,
    trimmer.WithSuffix("")))
// Output: "Short"

// Note: Using struct options with empty suffix will use default "..."
options = &trimmer.SmartTrimOptions{
    Suffix: "",  // Will be treated as default "..."
}
fmt.Println(trimmer.SmartTrim("Short text", 8, options))
// Output: "Short..."
```

## 📋 API

### SmartTrim Function

```go
func SmartTrim(text string, maxLength int, options *SmartTrimOptions, functionalOptions ...Option) string
```

#### Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `text` | `string` | The text to trim |
| `maxLength` | `int` | Maximum character length (including suffix) |
| `options` | `*SmartTrimOptions` | Optional configuration (can be nil) |
| `functionalOptions` | `...Option` | Functional options (variadic) |

#### SmartTrimOptions Struct

```go
type SmartTrimOptions struct {
    Suffix              string
    PreserveWholeWords  bool
    PreservePunctuation bool
}
```

| Field | Type | Default | Description |
|--------|------|---------|-------------|
| `Suffix` | `string` | `"..."` | Text to append after trimming |
| `PreserveWholeWords` | `bool` | `true` | When true, keeps words intact; when false, cuts mid-word |
| `PreservePunctuation` | `bool` | `true` | Keep or remove trailing punctuation |

#### Functional Options

```go
func WithSuffix(suffix string) Option
func WithPreserveWholeWords(preserve bool) Option
func WithPreservePunctuation(preserve bool) Option
```

## ✨ Key Features

- **Flexible word handling**: Can either preserve whole words (default) or cut mid-word when needed
- **Customizable suffix**: Use any string as the truncation indicator
- **Punctuation awareness**: Intelligently handles trailing punctuation
- **Thoroughly tested**: Handles all edge cases reliably
- **Zero dependencies**: Lightweight and efficient
- **Dual API**: Both struct-based and functional options for flexibility
- **High performance**: Optimized for speed with minimal allocations

## 🧩 Behavior Details

- Returns the original string if it's already shorter than `maxLength`
- Truncates the suffix if `maxLength` is smaller than the suffix length
- When `preserveWholeWords` is enabled with very long words, returns only the suffix
- Provides clear panic messages for invalid inputs
- When using struct options, omitting the Suffix field or setting it to an empty string will use the default suffix ("...")
- To explicitly set an empty suffix, use the functional option `WithSuffix("")`

## 📦 Why Use This Library?

While text trimming seems simple, handling all the edge cases correctly can be tricky:

- What if there are no spaces?
- How should punctuation be handled?
- What if the max length is very small?

This library has been thoroughly tested against dozens of edge cases so you don't have to worry about them.

## 🔍 Performance

The library is optimized for performance with:

- Minimal memory allocations
- Efficient string handling
- Fast punctuation detection
- Pre-allocated string builders

## 📄 License

MIT