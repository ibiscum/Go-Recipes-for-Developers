package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	input, err := os.Open("romeo-and-juliet.txt")
	if err != nil {
		panic(err)
	}
	// Search for lines containing weak
	cmd := exec.Command("grep", "weak")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	// Send the output to stdout
	cmd.Stdout = os.Stdout
	// Start the command
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	// Send the file to the command
	io.Copy(stdin, input)
	//  Close the stdin so the command knows it's done
	stdin.Close()
	// Wait until the program ends
	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}
