package casify

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func isLowerOrNumber(ch rune) bool {
	return unicode.IsLower(ch) || unicode.IsNumber(ch)
}

// Splits string into words on non-alphanumeric characters and proper camel case changes.
// Applies transform to each word, specified by callback, and joins the resulting words using separator into a single string.
// This function properly handles acronyms, e.g. TCPOrUDPSocket is 4 words: "TCP", "Or", "UDP", "Socket".
// Numbers are considered lowercase letters for the purpose of this function.
// Generally a number should not be the first character in a word, because it trips up the parser.
func Convert(s string, sep string, cb func(string) string) string {
	words := make([]string, 0)
	prevWordStart := 0

	processNextWord := func (wordEnd int) {
		word := s[prevWordStart:wordEnd]
		if len(word) > 0 {
			words = append(words, cb(word))
		}
		prevWordStart = wordEnd
	}

	for i, ch := range s {
		nextChOffset := i + utf8.RuneLen(ch)
		nextCh, _ := utf8.DecodeRuneInString(s[nextChOffset:])
		if !unicode.IsLetter(ch) && !unicode.IsNumber(ch) {
			processNextWord(i)
			prevWordStart = nextChOffset
		} else if unicode.IsUpper(ch) && isLowerOrNumber(nextCh) {
			processNextWord(i)
		} else if isLowerOrNumber(ch) && unicode.IsUpper(nextCh) {
			processNextWord(nextChOffset)
		}
	}
	processNextWord(len(s))
	return strings.Join(words, sep)
}

// Kinda "reverse" of what strings.Title does.
// Makes first rune in each alphanumeric word a lowercase letter.
// For a list of differences with strings.Title refer to the source code.
func Untitle(s string) string {
	prev := rune(0)
	return strings.Map(
		func(r rune) rune {
			if !unicode.IsLetter(prev) && !unicode.IsDigit(prev) {
				prev = r
				return unicode.ToLower(r)
			}
			prev = r
			return r
		},
		s)
}

// A shorthand for Convert(s, "_", strings.ToLower).
func SnakeCase(s string) string {
	return Convert(s, "_", strings.ToLower)
}

// A shorthand for Convert(s, "_", strings.ToUpper).
func UpperSnakeCase(s string) string {
	return Convert(s, "_", strings.ToUpper)
}

// A shorthand for Convert(s, "-", strings.ToLower).
func KebabCase(s string) string {
	return Convert(s, "-", strings.ToLower)
}

// A shorthand for Convert(s, "-", strings.ToUpper).
func UpperKebabCase(s string) string {
	return Convert(s, "-", strings.ToUpper)
}

// A shorthand for Untitle(Convert(s, "-", strings.Title ∘ strings.ToLower)), here ∘ is a functional composition operator.
func CamelCase(s string) string {
	return Untitle(Convert(s, "", func (s string) string {
		return strings.Title(strings.ToLower(s))
	}))
}

// A shorthand for Convert(s, "-", strings.Title ∘ strings.ToLower), here ∘ is a functional composition operator.
func UpperCamelCase(s string) string {
	return Convert(s, "", func (s string) string {
		return strings.Title(strings.ToLower(s))
	})
}
