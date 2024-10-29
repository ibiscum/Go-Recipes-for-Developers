package main

import (
	"fmt"
	"regexp"
)

var integerRegexp = regexp.MustCompile("^[0-9]+$")

func main() {
	fmt.Println(integerRegexp.MatchString("123"))   // true
	fmt.Println(integerRegexp.MatchString(" 123 ")) // false
}
