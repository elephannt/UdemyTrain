package main

import (
	"fmt"
	"regexp"
)

func main() {
	value := "marco"

	// First compile the delimiter expression.
	re := regexp.MustCompile(`^[a-z]+$`)

	// Split based on pattern.
	// ... Second argument means "no limit."
	result := re.Split(value, -1)

	for i := range(result) {
		fmt.Println(result[i])
	}
}