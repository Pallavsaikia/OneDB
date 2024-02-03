package libs

import (
	"unicode"
)

func ContainsSpace(s string) bool {
	for _, char := range s {
		if unicode.IsSpace(char) {
			return true
		}
	}
	return false
}

func ContainsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func ContainsSpecialCharacters(s string) bool {
	allowedCharacters := map[rune]bool{
		'_': true,
	}
	for _, char := range s {
		if !unicode.IsLetter(char) && !allowedCharacters[char] {
			return true
		}
	}
	return false
}

