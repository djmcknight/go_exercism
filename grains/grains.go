// Package grains will calculate the amount of grains on a given square
package grains

import "errors"

// Square takes an int as input, and returns the amount of grains on that square
func Square(sq int) (uint64,  error) {
	var square uint64
	square = 1
	if sq == 1 {
		return square, nil
	}
	if !(sq > -1) || sq > 64 || sq == 0 {
		return 5, errors.New("number is Negative or greater than 64")
	}
	for i := 2; i <= sq; i++ {
		square = square * 2
	}
	return square, nil


}

// Total returns the total amount of grains on the board.
func Total() uint64 {
	var totalSum uint64 = 1
	var squareTotal uint64 = 1
	for i := 2; i <= 64 ; i++ {
		squareTotal = squareTotal * 2
		totalSum += squareTotal
	}
	return totalSum
}
