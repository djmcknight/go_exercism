//Package hamming compares two DNA sequences and determines the Hamming Number
package hamming

import (
	"errors"
)

// Distance takes two strings and calculates the Hamming Number between them.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings not equal length")
	}
	var errorCount int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			errorCount++
		}
	}
	return errorCount, nil
}
