package cppstring

import "github.com/West-Pavilion/cppstring/v2/helpers"

type Cppstring struct {
	rawString string
}

func (cstring *Cppstring) Append(otherString *Cppstring) Cppstring {
	return Cppstring{string(append([]rune(cstring.rawString), []rune(otherString.rawString)...))}
}

func Reverse(raw string) string {
	return helpers.Reverse(raw)
}
