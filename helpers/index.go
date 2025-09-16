// Package helpers contains the internal implementation of string manipulation functions.
package helpers

// Reverse reverses a string by converting it to a slice of runes and swapping characters.
// This approach correctly handles Unicode characters, including multi-byte characters
// like Chinese characters and emojis.
func Reverse(string2 string) string {
	runes := []rune(string2)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
