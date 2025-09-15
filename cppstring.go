package cppstring

import "github.com/West-Pavilion/cppstring/v2/helpers"

type cppstring struct {
	rawString string
}

func (cstring *cppstring) append(otherString *cppstring) cppstring {
	return cppstring{string(append([]rune(cstring.rawString), []rune(otherString.rawString)...))}
}

func Reverse(raw string) string {
	return helpers.Reverse(raw)
}
