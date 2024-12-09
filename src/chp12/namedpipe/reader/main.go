package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the named pipe for reading
	pipe, err := os.Open("../pipe")
	if err != nil {
		panic(err)
	}
	defer pipe.Close()

	// Read data until it is closed
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
