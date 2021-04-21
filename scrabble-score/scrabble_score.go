// Package scrabble calculates your score when playing scrabble
package scrabble

import (
	"strings"
)

// Score calculates your score by adding up the individual letter value
func Score(word string) int {
	word = strings.ToUpper(word)
	var score int
	for i := 0; i < len(word); i++{
		switch word[i] {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		default:
			return 0
		}
	}
	return score
}
