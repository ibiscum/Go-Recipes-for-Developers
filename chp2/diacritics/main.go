// Based on the blog post https://go.dev/blog/normalization
package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.
		NFC)
	rd := transform.NewReader(strings.NewReader("Montréal"), t)
	str, _ := io.ReadAll(rd)
	fmt.Println(string(str))
}
