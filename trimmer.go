package trimmer

import (
	"fmt"
	"strings"
)

type SmartTrimOptions struct {
	Suffix              string
	PreserveWholeWords  bool
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

// resolveOptions combines struct options and functional options
func resolveOptions(opts *SmartTrimOptions, options ...Option) SmartTrimOptions {
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

	return resolvedOpts
}

func calculateContentLength(maxLength, suffixLen int) int {
	contentLength := maxLength - suffixLen

	if contentLength <= 0 {
		return 0
	}

	return contentLength
}

func handleSmallMaxLength(maxLength int, suffix string) string {
	if maxLength <= 0 {
		return ""
	}
	return suffix[:maxLength]
}

func trimAtWordBoundary(text, trimmedText string, contentLength, textLen int) string {
	if contentLength < textLen && text[contentLength] != ' ' {
		lastSpaceIndex := strings.LastIndex(trimmedText, " ")
		if lastSpaceIndex != -1 {
			return trimmedText[:lastSpaceIndex]
		}
		return ""
	}
	return trimmedText
}

// removes trailing punctuation from text
func removePunctuation(text string) string {
	for len(text) > 0 && isPunctuation(rune(text[len(text)-1])) {
		text = text[:len(text)-1]
	}
	return text
}

func SmartTrim(text string, maxLength int, opts *SmartTrimOptions, options ...Option) string {
	if maxLength < 0 {
		panic(fmt.Sprintf("SmartTrim: maxLength must be a non-negative integer, got %d", maxLength))
	}

	textLen := len(text)
	if textLen <= maxLength {
		return text
	}

	resolvedOpts := resolveOptions(opts, options...)
	suffixLen := len(resolvedOpts.Suffix)
	contentLength := calculateContentLength(maxLength, suffixLen)

	if contentLength <= 0 {
		return handleSmallMaxLength(maxLength, resolvedOpts.Suffix)
	}

	trimmedText := text[:contentLength]

	if resolvedOpts.PreserveWholeWords && contentLength < textLen {
		trimmedText = trimAtWordBoundary(text, trimmedText, contentLength, textLen)
	}

	trimmedText = strings.TrimRight(trimmedText, " \t\n\r")

	if !resolvedOpts.PreservePunctuation && len(trimmedText) > 0 {
		trimmedText = removePunctuation(trimmedText)
	}

	var sb strings.Builder
	sb.Grow(len(trimmedText) + suffixLen) // Pre-allocate exact capacity
	sb.WriteString(trimmedText)
	sb.WriteString(resolvedOpts.Suffix)
	return sb.String()
}
