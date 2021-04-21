// Package leap determines if the input year is a leap year or not
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(y int) bool {
	if y%4 == 0 && y%400 == 0 {
		return true
	} else if y%4 == 0 && y%100 == 0 {
		return false
	} else if y%4 == 0 {
		return true
	} else {
		return false
	}
}
