package trimmer

import (
	"fmt"
	"strings"
)

type SmartTrimOptions struct {
	Suffix string
	PreserveWholeWords bool
	PreservePunctuation bool
}

type Option func(*SmartTrimOptions)

func WithSuffix(suffix string) Option {
	return func(o *SmartTrimOptions) {
		o.Suffix = suffix
	}
}

func WithPreserveWholeWords(preserve bool) Option {
	return func(o *SmartTrimOptions) {
		o.PreserveWholeWords = preserve
	}
}

func WithPreservePunctuation(preserve bool) Option {
	return func(o *SmartTrimOptions) {
		o.PreservePunctuation = preserve
	}
}

func DefaultOptions() SmartTrimOptions {
	return SmartTrimOptions{
		Suffix:              "...",
		PreserveWholeWords:  true,
		PreservePunctuation: true,
	}
}

func isPunctuation(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ';', ':', '\'', '"', ')', ']', '}', '…':
		return true
	default:
		return false
	}
}

func SmartTrim(text string, maxLength int, opts *SmartTrimOptions, options ...Option) string {
	if maxLength < 0 {
		panic(fmt.Sprintf("SmartTrim: maxLength must be a non-negative integer, got %d", maxLength))
	}

	textLen := len(text)
	if textLen <= maxLength {
		return text
	}

	resolvedOpts := DefaultOptions()

	if opts != nil {
		if opts.Suffix != "" {
			resolvedOpts.Suffix = opts.Suffix
		}
		resolvedOpts.PreserveWholeWords = opts.PreserveWholeWords
		resolvedOpts.PreservePunctuation = opts.PreservePunctuation
	}

	for _, option := range options {
		option(&resolvedOpts)
	}

	suffixLen := len(resolvedOpts.Suffix)
	contentLength := maxLength - suffixLen

	// Handle case where maxLength is too small for anything but (part of) the suffix
	if contentLength <= 0 {
		if maxLength <= 0 {
			return ""
		}
		return resolvedOpts.Suffix[:maxLength]
	}

	// Ensure not go out of bounds
	if contentLength > textLen {
		contentLength = textLen
	}

	trimmedText := text[:contentLength]

	if resolvedOpts.PreserveWholeWords && contentLength < textLen {
		// Check if we're in the middle of a word
		if contentLength < textLen && text[contentLength] != ' ' {
			lastSpaceIndex := strings.LastIndex(trimmedText, " ")
			if lastSpaceIndex != -1 {
				trimmedText = trimmedText[:lastSpaceIndex]
			} else {
				trimmedText = ""
			}
		}
	}

	trimmedText = strings.TrimRight(trimmedText, " \t\n\r")

	if !resolvedOpts.PreservePunctuation && len(trimmedText) > 0 {
		for len(trimmedText) > 0 && isPunctuation(rune(trimmedText[len(trimmedText)-1])) {
			trimmedText = trimmedText[:len(trimmedText)-1]
		}
	}

	var sb strings.Builder
	sb.Grow(len(trimmedText) + suffixLen) // Pre-allocate exact capacity
	sb.WriteString(trimmedText)
	sb.WriteString(resolvedOpts.Suffix)
	return sb.String()
}
