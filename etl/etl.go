//Package etl transforms your old scoring structure and turns it into the more efficient scoring system.
package etl

import "strings"

// Transform takes your old scoring system and transforms it to a per letter system
func Transform(input map[int][]string) map[string]int {
	etlMap := make(map[string]int)
	for i, k := range input {
		for _, r := range k {
			etlMap[strings.ToLower(r)] = i
		}
	}
	return etlMap
}
