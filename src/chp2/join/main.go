package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	strs := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(strs, " "))
	fmt.Println(strings.Join(strs, ""))
	fmt.Println(path.Join(strs...))
	fmt.Println(filepath.Join(strs...))

	strs2 := []string{"/foo", "//bar", "baz"}
	fmt.Println(strings.Join(strs2, " "))
	fmt.Println(strings.Join(strs2, ""))
	fmt.Println(path.Join(strs2...))
	fmt.Println(filepath.Join(strs2...))
}
