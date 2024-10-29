package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("Break-------", "-"))
	fmt.Println(strings.TrimRight("Break with spaces-- -- --", "- "))
	fmt.Println(strings.TrimSuffix("file.txt", ".txt"))
	fmt.Println(strings.TrimLeft(" \t   Indented text", " \t"))
	fmt.Println(strings.TrimSpace(" \t \n  Indented text  \n\t"))
}
