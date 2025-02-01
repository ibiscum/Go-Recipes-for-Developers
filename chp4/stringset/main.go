package main

import (
	"fmt"
	"strings"
)

// Remove duplicate inputs from the input, preserving order
func DedupOrdered(input []string) []string {
	set := make(map[string]struct{})
	output := make([]string, 0, len(input))
	for _, in := range input {
		if _, exists := set[in]; exists {
			continue
		}
		output = append(output, in)
		set[in] = struct{}{}
	}
	return output
}

func main() {
	str := `this function removes the duplicate words in the input using a map used as a string set`
	fmt.Println("Input:", str)
	fmt.Println("Output:", DedupOrdered(strings.Split(str, " ")))
}
