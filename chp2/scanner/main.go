package main

import (
	"bufio"
	"fmt"
	"strings"
)

const input = `This is a string
that has 3
lines.`

func main() {
	lineScanner := bufio.NewScanner(strings.NewReader(input))
	line := 0
	for lineScanner.Scan() {
		text := lineScanner.Text()
		line++
		fmt.Printf("Line %d: %s\n", line, text)
	}
	if err := lineScanner.Err(); err != nil {
		panic(err)
	}
	wordScanner := bufio.NewScanner(strings.NewReader(input))
	wordScanner.Split(bufio.ScanWords)
	word := 0
	for wordScanner.Scan() {
		text := wordScanner.Text()
		word++
		fmt.Printf("word %d: %s\n", word, text)
	}
	if err := wordScanner.Err(); err != nil {
		panic(err)
	}
}
