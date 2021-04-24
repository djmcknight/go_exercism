package reverse

import (
	"unicode/utf8"
)

func Reverse(word string) string {
	 wordB := []byte(word)
	 var drow string
	for len(wordB) > 0 {
		r, size := utf8.DecodeLastRune(wordB)
		drow += string(r)
		wordB = wordB[:len(wordB)-size]
	}
	return drow
}