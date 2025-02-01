package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// Write to the file
	data := []byte("Hello, World!")
	count, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", count)
	fsize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File size: %d\n", fsize)

	// Set file size to 1000
	err = file.Truncate(1000)
	if err != nil {
		panic(err)
	}
	fsize, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File size: %d\n", fsize)

	// Move to offset `count`
	_, err = file.Seek(int64(count), io.SeekStart)
	if err != nil {
		panic(err)
	}
	// Write to the file
	_, err = file.Write([]byte("\n"))
	if err != nil {
		panic(err)
	}
	// Seek beyond file end
	_, err = file.Seek(2000, io.SeekStart)
	if err != nil {
		panic(err)
	}
	// Write to the file
	_, err = file.Write([]byte{0})
	if err != nil {
		panic(err)
	}
	fsize, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File size: %d\n", fsize)

	_, err = file.WriteAt([]byte("Hello, World!"), 10)
	if err != nil {
		panic(err)
	}
	offset, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		panic(err)
	}
	fmt.Printf("After WriteAt position is: %d\n", offset)

	buffer := make([]byte, 5)
	file.ReadAt(buffer, 10)
	fmt.Println("ReadAt 10:", string(buffer))

	file.Close()

	// Truncate open file
	file, err = os.OpenFile("test.txt", os.O_RDWR|os.O_TRUNC, 0o644)
	if err != nil {
		panic(err)
	}
	file.Close()
}
