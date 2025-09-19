// Package cppstring provides C++-style string manipulation functions in Go.
// Despite the name, this is a pure Go implementation designed to offer
// efficient and Unicode-aware string operations.
package cppstring

import "github.com/West-Pavilion/cppstring/helpers"

var Version = "NoClass"

// Reverse returns a new string with the characters of the input string in reverse order.
// It properly handles Unicode characters including Chinese characters and emojis.
func Reverse(raw string) string {
	return helpers.Reverse(raw)
}
