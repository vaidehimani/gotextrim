package trimmer

import (
	"testing"
)

func TestSmartTrim(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		maxLen   int
		options  *SmartTrimOptions
		expected string
	}{
		{
			name:     "Text shorter than max length",
			text:     "Hello world",
			maxLen:   20,
			options:  nil,
			expected: "Hello world",
		},
		{
			name:     "Basic trimming with default options",
			text:     "The quick brown fox jumps over the lazy dog",
			maxLen:   20,
			options:  nil,
			expected: "The quick brown...",
		},
		{
			name:   "Custom suffix",
			text:   "The quick brown fox jumps over the lazy dog",
			maxLen: 25,
			options: &SmartTrimOptions{
				Suffix:              " [more]",
				PreserveWholeWords:  true,
				PreservePunctuation: true,
			},
			expected: "The quick brown [more]",
		},
		{
			name:   "Don't preserve whole words",
			text:   "The quick brown fox jumps over the lazy dog",
			maxLen: 20,
			options: &SmartTrimOptions{
				Suffix:              "...",
				PreserveWholeWords:  false,
				PreservePunctuation: true,
			},
			expected: "The quick brown f...",
		},
		{
			name:   "Don't preserve punctuation",
			text:   "The quick brown fox, jumps over the lazy dog",
			maxLen: 24,
			options: &SmartTrimOptions{
				Suffix:              "...",
				PreserveWholeWords:  true,
				PreservePunctuation: false,
			},
			expected: "The quick brown fox...",
		},
		{
			name:   "Preserve punctuation",
			text:   "The quick brown fox, jumps over the lazy dog",
			maxLen: 24,
			options: &SmartTrimOptions{
				Suffix:              "...",
				PreserveWholeWords:  true,
				PreservePunctuation: true,
			},
			expected: "The quick brown fox,...",
		},
		{
			name:     "Very short max length",
			text:     "Hello world",
			maxLen:   2,
			options:  nil,
			expected: "..",
		},
		{
			name:   "Text with trailing punctuation",
			text:   "Hello world!",
			maxLen: 8,
			options: &SmartTrimOptions{
				PreservePunctuation: false,
			},
			expected: "Hello...",
		},
		{
			name:   "Text with trailing punctuation preserved",
			text:   "Hello world!",
			maxLen: 8,
			options: &SmartTrimOptions{
				PreservePunctuation: true,
			},
			expected: "Hello...",
		},
		{
			name:     "Empty string",
			text:     "",
			maxLen:   5,
			options:  nil,
			expected: "",
		},
		{
			name:     "Mid-word break detection",
			text:     "ThisIsAVeryLongWordWithoutSpaces",
			maxLen:   10,
			options:  nil,
			expected: "...",
		},
		{
			name:     "Cut exactly at word boundary",
			text:     "One Two Three Four",
			maxLen:   9,
			options:  nil,
			expected: "One...",
		},
		{
			name:   "Multiple punctuation characters",
			text:   "Hello, world!?",
			maxLen: 10,
			options: &SmartTrimOptions{
				PreservePunctuation: false,
			},
			expected: "Hello...",
		},
		{
			name:     "Max length greater than text length",
			text:     "Short text",
			maxLen:   100,
			options:  nil,
			expected: "Short text",
		},
		{
			name:     "Handle extra space",
			text:     "Short text  check",
			maxLen:   14,
			options:  nil,
			expected: "Short text...",
		},
		{
			name:   "Handle no suffix using struct",
			text:   "Short text",
			maxLen: 6,
			options: &SmartTrimOptions{
				Suffix:              "",
				PreserveWholeWords:  true,
				PreservePunctuation: true,
			},
			expected: "...",
		},
		{
			name:   "Handle suffix longer than maxlength ",
			text:   "Short text",
			maxLen: 3,
			options: &SmartTrimOptions{
				Suffix:              "more",
				PreserveWholeWords:  true,
				PreservePunctuation: true,
			},
			expected: "mor",
		},
		{
			name:     "trims input containing only whitespace",
			text:     "          ",
			maxLen:   3,
			options:  nil,
			expected: "...",
		},
		{
			name:   "handles string made of only punctuation",
			text:   "!!!???...",
			maxLen: 5,
			options: &SmartTrimOptions{
				Suffix:              "...",
				PreserveWholeWords:  true,
				PreservePunctuation: false,
			},
			expected: "...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SmartTrim(tt.text, tt.maxLen, tt.options)
			if result != tt.expected {
				t.Errorf("SmartTrim() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFunctionalOptions(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog"

	// Test with functional options
	result := SmartTrim(
		text,
		22,
		nil,
		WithSuffix(" [read more]"),
		WithPreserveWholeWords(true),
	)

	expected := "The quick [read more]"
	if result != expected {
		t.Errorf("SmartTrim with functional options = %q, want %q", result, expected)
	}

	// Test mixing struct options and functional options (functional should override)
	baseOpts := &SmartTrimOptions{
		Suffix:              "...",
		PreserveWholeWords:  true,
		PreservePunctuation: true,
	}

	result = SmartTrim(
		text,
		20,
		baseOpts,
		WithSuffix(" [override]"),
	)

	expected = "The quick [override]"
	if result != expected {
		t.Errorf("SmartTrim with mixed options = %q, want %q", result, expected)
	}

	// Test functional option with empty suffix
	result = SmartTrim(
		"Short text",
		6,
		nil,
		WithSuffix(""),
	)

	expected = "Short"
	if result != expected {
		t.Errorf("SmartTrim with empty suffix functional option = %q, want %q", result, expected)
	}
}

func TestDefaultOptions(t *testing.T) {
	opts := DefaultOptions()

	if opts.Suffix != "..." {
		t.Errorf("DefaultOptions().Suffix = %q, want %q", opts.Suffix, "...")
	}

	if !opts.PreserveWholeWords {
		t.Errorf("DefaultOptions().PreserveWholeWords = %v, want %v", opts.PreserveWholeWords, true)
	}

	if !opts.PreservePunctuation {
		t.Errorf("DefaultOptions().PreservePunctuation = %v, want %v", opts.PreservePunctuation, true)
	}
}

func TestPanicOnNegativeLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for negative maxLength, but no panic occurred")
		}
	}()

	SmartTrim("Test", -1, nil)
}

func TestIsPunctuation(t *testing.T) {
	punctuations := []rune{'.', ',', '!', '?', ';', ':', '\'', '"', ')', ']', '}', '…'}
	nonPunctuations := []rune{'a', 'B', '1', ' ', '\t', '(', '[', '{'}

	for _, r := range punctuations {
		if !isPunctuation(r) {
			t.Errorf("isPunctuation(%q) = false, want true", r)
		}
	}

	for _, r := range nonPunctuations {
		if isPunctuation(r) {
			t.Errorf("isPunctuation(%q) = true, want false", r)
		}
	}
}
