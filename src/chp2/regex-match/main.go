package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[0-9]+`)
	fmt.Println(re.FindAllString("This regular expression find numbers, like 1, 100, 500, etc.", -1))
}
