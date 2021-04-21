// Package twofer returns a string based on the name parameter provided.
package twofer

import (
	"fmt"
)

// ShareWith returns string based on name, and default if name not matched in switch.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
