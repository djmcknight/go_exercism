// Package raindrops will return a string, or number, based on the modulus of the number.
package raindrops

import (
	"strconv"
)

// Convert will return a string based on the modulus of the number provided.
func Convert(number int) string {
	var response string
	if number%3 == 0 {
		response += "Pling"
	}
	if number%5 == 0 {
		response += "Plang"
	}
	if number%7 == 0 {
		response += "Plong"
	}
	if response == "" {
		response = strconv.Itoa(number)
	}
	return response
}
