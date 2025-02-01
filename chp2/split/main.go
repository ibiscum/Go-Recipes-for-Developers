package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("a,b,c,d", ","))
	fmt.Println(strings.Split("a, b, c, d", ","))
	fmt.Println(strings.Fields("a    b   c  d  "))
	for i, x := range strings.Split("a    b   c  d  ", " ") {
		if i > 0 {
			fmt.Printf(",  ")
		}
		fmt.Printf(`"%s"`, x)
	}
	fmt.Println()
	for i, x := range strings.Split("a---b---c--d--", "-") {
		if i > 0 {
			fmt.Printf(",  ")
		}
		fmt.Printf(`"%s"`, x)
	}
	fmt.Println()
}
