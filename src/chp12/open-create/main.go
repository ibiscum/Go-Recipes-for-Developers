package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func open() {
	// Open a file read-only
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	// Make sure it is closed
	defer file.Close()

	// Read from the file
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			panic(err)
		}
	}
	fmt.Printf("Read %d bytes: %s\n", count, string(data))
	// Attempt to write to a read-only file
	_, err = file.Write([]byte("Hello, World!"))
	fmt.Printf("Attempt to write returns: %T %v\n", err, err)
}

func main() {
	open()
}
