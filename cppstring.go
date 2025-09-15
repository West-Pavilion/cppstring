package cppstring

import "github.com/West-Pavilion/cppstring/v2/helpers"

type Cppstring struct {
	RawString string
}

func (cstring *Cppstring) Append(otherString *string) Cppstring {
	cstring.RawString = string(append([]rune(cstring.RawString), []rune(*otherString)...))
	return *cstring
}

func (cstring *Cppstring) AppendCppString(otherString *Cppstring) Cppstring {
	return Cppstring{string(append([]rune(cstring.RawString), []rune(otherString.RawString)...))}
}

func Reverse(raw string) string {
	return helpers.Reverse(raw)
}
