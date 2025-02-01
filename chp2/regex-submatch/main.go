package main

import (
	"fmt"
	"regexp"
)

func main() {
	// This regular expression allows extra spaces before and after the
	// property name and value using \s*
	re := regexp.MustCompile(`^\s*(\w+)\s*=\s*(\w+)\s*$`)
	result := re.FindStringSubmatch(`  property = 12 `)
	fmt.Printf("Key: %s value: %s\n", result[1], result[2])
	result = re.FindStringSubmatch(`x = y `)
	fmt.Printf("Key: %s value: %s\n", result[1], result[2])
}
