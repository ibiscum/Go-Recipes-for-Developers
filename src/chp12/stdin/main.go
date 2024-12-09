package main

import (
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
	// Feed the file to the standard input
	cmd.Stdin = input
	// Redirect the output to stdout
	cmd.Stdout = os.Stdout
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}
