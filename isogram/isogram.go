// Package isogram will calculate if a word is an isogram
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word is an isogram
func IsIsogram(word string) bool {
	word = strings.ToLower(strings.Trim(word, ""))
	isoWord := make(map[rune]bool)
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		if _ , ok := isoWord[r]; ok {
			return false
		}
		isoWord[r] = true
	}
	return true
}
