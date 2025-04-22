package trimmer_test

import (
	"fmt"

	trimmer "github.com/vaidehimani/gotextrim"
)

func ExampleSmartTrim() {
	// Basic usage with default options
	text := "The quick brown fox jumps over the lazy dog"
	trimmed := trimmer.SmartTrim(text, 20, nil)
	fmt.Println(trimmed)

	// Custom suffix
	customOptions := &trimmer.SmartTrimOptions{
		Suffix: " [more]",
	}
	trimmedCustom := trimmer.SmartTrim(text, 20, customOptions)
	fmt.Println(trimmedCustom)

	// Don't preserve whole words
	noWholeWordsOptions := &trimmer.SmartTrimOptions{
		PreserveWholeWords: false,
	}
	trimmedNoWhole := trimmer.SmartTrim(text, 20, noWholeWordsOptions)
	fmt.Println(trimmedNoWhole)

	// Don't preserve trailing punctuation
	text = "Hello, world!"
	noPunctOptions := &trimmer.SmartTrimOptions{
		PreservePunctuation: false,
	}
	trimmedNoPunct := trimmer.SmartTrim(text, 9, noPunctOptions)
	fmt.Println(trimmedNoPunct)

	// Using functional options pattern (recommended)
	fmt.Println(trimmer.SmartTrim("This is a long sentence", 18, nil,
		trimmer.WithSuffix(" [read more]"),
		trimmer.WithPreserveWholeWords(true)))

	// Set empty suffix with functional option
	fmt.Println(trimmer.SmartTrim("Short text", 6, nil,
		trimmer.WithSuffix("")))

	// Note: Using struct options with empty suffix will use default "..."
	options := &trimmer.SmartTrimOptions{
		Suffix: "", // Will be treated as default "..."
	}
	fmt.Println(trimmer.SmartTrim("Short text", 8, options))

	// Output:
	// The quick brown...
	// The quick bro [more]
	// The quick brown f...
	// Hello...
	// This [read more]
	// Short
	// Short...
}
