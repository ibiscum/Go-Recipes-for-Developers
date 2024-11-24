package main

import (
	"fmt"
	"os"

	"golang.org/x/text/encoding/ianaindex"
)

func main() {
	enc, err := ianaindex.MIME.Encoding("US-ASCII")
	if err != nil {
		panic(err)
	}
	b, err := os.ReadFile("ascii.txt")
	if err != nil {
		panic(err)
	}
	decoder := enc.NewDecoder()
	encoded, err := decoder.Bytes(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))
}
