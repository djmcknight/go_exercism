// Package luhn will verify if a provided string is a valid luhn number
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)


// Valid performs a check on the string, returning a bool if the input is a luhn number or not
func Valid(input string) bool {
	input = strings.Join(strings.Fields(input),"")
	var luhnVal int

	if len(input) <= 1 {
		return false
	}
	var everyOther int
	for i := len(input) -1;  i >= 0; i--{

		if !unicode.IsNumber(rune(input[i])) {
			return false
		}
		x, _ := strconv.Atoi(string(input[i]))
		if everyOther % 2 != 0 {
			x = x * 2
			if x > 9 {
				x -= 9
			}
		}
		luhnVal += x
		everyOther++


	}
	if luhnVal % 10 != 0 {
		return false
	}
	return true
}