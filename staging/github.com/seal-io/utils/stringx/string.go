package stringx

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Join concatenates the elements of strs to create a single string.
func Join[T ~string](sep string, strs ...T) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return string(strs[0])
	}

	n := len(sep) * (len(strs) - 1)
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}

	var b strings.Builder

	b.Grow(n)
	b.WriteString(string(strs[0]))

	for i := range strs[1:] {
		b.WriteString(sep)
		b.WriteString(string(strs[i+1]))
	}

	return b.String()
}

// Strings converts a slice of fmt.Stringer to a slice of string.
func Strings[T fmt.Stringer](v []T) []string {
	if len(v) == 0 {
		return nil
	}

	s := make([]string, len(v))
	for i := range v {
		s[i] = v[i].String()
	}

	return s
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// IsSpace reports whether the rune is a space character.
func IsSpace(r rune) bool {
	if r > utf8.RuneSelf {
		return unicode.IsSpace(r)
	}
	return asciiSpace[r] == 1
}

// TrimAllSpace removes space chars from the given string.
func TrimAllSpace(s string) string {
	if len(s) == 0 {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if IsSpace(r) {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// ReplaceFunc returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
func ReplaceFunc(s string, rp func(rune) rune) string {
	if len(s) == 0 || rp == nil {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		b.WriteRune(rp(r))
	}
	return b.String()
}
