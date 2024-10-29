package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Find numbers, capture the first digit
	re := regexp.MustCompile(`([0-9])[0-9]*`)
	fmt.Println(re.ReplaceAllString("This example replaces numbers  with 'x': 1, 100, 500.", "x"))
	fmt.Println(re.ReplaceAllString("This example replaces all numbers with their first digits: 1, 100, 500.", "${1}"))
}
