package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Open the named pipe
	pipe, err := os.OpenFile("../pipe", os.O_WRONLY, fs.ModeNamedPipe|0660)
	if err != nil {
		panic(err)
	}
	defer pipe.Close()

	// Generate data and write to pipe
	for i := 0; i < 20; i++ {
		fmt.Fprintf(pipe, "Data %d\n", i)
	}
}
