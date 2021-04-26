package grains

import (
	"errors"
)

// Square takes an int as input, and returns the amount of grains on that square
func Square(sq int) (uint64, error) {
	if sq < 1 || sq > 64 {
		return 0, errors.New("number is Negative or greater than 64")
	}
	return 1 << (sq - 1), nil
}

// Total returns the total amount of grains on the board.
func Total() uint64 {
	//18446744073709551615
	return 1<<64 - 1
}
