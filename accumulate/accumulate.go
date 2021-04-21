// Package accumulate modifies an input string based on the converter provided.
package accumulate

// Accumulate manipulates an input string based on the type of converter provided
func Accumulate(inputWords []string, converter func(string) string) []string {
	for i := 0; i < len(inputWords); i++ {
		inputWords[i] = converter(inputWords[i])
	}
	return inputWords
}
