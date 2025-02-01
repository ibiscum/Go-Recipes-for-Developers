package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	words := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(words, " "))
	// foo bar baz
	fmt.Println(strings.Join(words, ""))
	// foobarbaz
	fmt.Println(path.Join(words...))
	// foo/bar/baz
	fmt.Println(filepath.Join(words...))
	// foo/bar/baz or foo\bar\baz, depending on the host system
	paths := []string{"/foo", "//bar", "baz"}
	fmt.Println(strings.Join(paths, " "))
	// /foo //bar baz
	fmt.Println(path.Join(paths...))
	// /foo/bar/baz
	fmt.Println(filepath.Join(paths...))
	// /foo/bar/baz or \foo\bar\baz depending on the host system}
}
