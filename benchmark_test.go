package trimmer

import (
	"strings"
	"testing"
)

var benchResult string // To prevent compiler optimizations

// Benchmark for different input sizes
func BenchmarkSmartTrim_InputSize(b *testing.B) {
	shortText := "Hello world"
	mediumText := "The quick brown fox jumps over the lazy dog. This is a medium-sized text for benchmarking."
	longText := strings.Repeat(mediumText, 10)
	veryLongText := strings.Repeat(mediumText, 100)

	texts := map[string]string{
		"Short":     shortText,
		"Medium":    mediumText,
		"Long":      longText,
		"Very Long": veryLongText,
	}

	for name, text := range texts {
		b.Run(name, func(b *testing.B) {
			var r string
			for i := 0; i < b.N; i++ {
				r = SmartTrim(text, 50, nil)
			}
			benchResult = r // Prevent compiler optimization
		})
	}
}

// Benchmark for different option combinations
func BenchmarkSmartTrim_Options(b *testing.B) {
	text := "The quick brown fox jumps over the lazy dog. This is a medium-sized text for benchmarking."

	b.Run("Default", func(b *testing.B) {
		var r string
		for i := 0; i < b.N; i++ {
			r = SmartTrim(text, 50, nil)
		}
		benchResult = r
	})

	b.Run("StructOptions", func(b *testing.B) {
		opts := &SmartTrimOptions{
			Suffix:              " [more]",
			PreserveWholeWords:  false,
			PreservePunctuation: false,
		}
		var r string
		for i := 0; i < b.N; i++ {
			r = SmartTrim(text, 50, opts)
		}
		benchResult = r
	})

	b.Run("FunctionalOptions", func(b *testing.B) {
		var r string
		for i := 0; i < b.N; i++ {
			r = SmartTrim(text, 50, nil,
				WithSuffix(" [more]"),
				WithPreserveWholeWords(false),
				WithPreservePunctuation(false),
			)
		}
		benchResult = r
	})
}

// Benchmark for different maxLength values
func BenchmarkSmartTrim_MaxLength(b *testing.B) {
	text := "The quick brown fox jumps over the lazy dog. This is a medium-sized text for benchmarking."

	lengths := map[string]int{
		"Very Short (5)":  5,   // Just suffix
		"Short (20)":      20,  // Part of the text
		"Medium (50)":     50,  // About half text
		"Long (100)":      100, // Full text
		"Very Long (500)": 500, // Much larger than text
	}

	for name, length := range lengths {
		b.Run(name, func(b *testing.B) {
			var r string
			for i := 0; i < b.N; i++ {
				r = SmartTrim(text, length, nil)
			}
			benchResult = r
		})
	}
}

// Benchmark comparing preservation settings
func BenchmarkSmartTrim_Preservation(b *testing.B) {
	text := "The quick brown fox jumps over the lazy dog. This is a medium-sized text for benchmarking!"

	b.Run("PreserveWordsAndPunct", func(b *testing.B) {
		var r string
		for i := 0; i < b.N; i++ {
			r = SmartTrim(text, 50, nil,
				WithPreserveWholeWords(true),
				WithPreservePunctuation(true),
			)
		}
		benchResult = r
	})

	b.Run("PreserveWordsNotPunct", func(b *testing.B) {
		var r string
		for i := 0; i < b.N; i++ {
			r = SmartTrim(text, 50, nil,
				WithPreserveWholeWords(true),
				WithPreservePunctuation(false),
			)
		}
		benchResult = r
	})

	b.Run("NoPreservation", func(b *testing.B) {
		var r string
		for i := 0; i < b.N; i++ {
			r = SmartTrim(text, 50, nil,
				WithPreserveWholeWords(false),
				WithPreservePunctuation(false),
			)
		}
		benchResult = r
	})
}
