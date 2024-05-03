package errors

import (
	"fmt"
	"regexp"
	"testing"
)

func TestSubnetNotFound(t *testing.T) {
	input := "subnet with id sub-f98ee-aec0-4662-a105-4d72b8c78b0a is not found"
	pattern := `subnet with id [^.]+ is not found`

	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	// Find the matching substring
	match := re.FindString(input)

	if match != "" {
		fmt.Printf("Match found: %s\n", match)
	} else {
		fmt.Println("No match found")
	}
}
