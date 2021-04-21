// Package diffsquares squares the sum or sums the square of N natural numbers
package diffsquares

// SquareOfSum sums 0 + n natural numbers and squares it
func SquareOfSum(n int) int {
	sqSum := 0
	for i := 0; i < n+1; i++ {
		sqSum += i
	}
	return sqSum * sqSum
}

// SumOfSquares squares o + n natural numbers then sums it
func SumOfSquares(n int) int {
	sumSq := 0
	for i := 0; i < n+1; i++ {
		sumSq += i * i
	}
	return sumSq
}

// Difference finds the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
