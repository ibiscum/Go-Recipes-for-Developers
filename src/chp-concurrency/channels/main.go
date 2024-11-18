package main

import "fmt"

var data []string = []string{
	"this",
	"example",
	"demonstrates",
	"using",
	"channels",
	"to",
	"send",
	"and",
	"receive",
	"data",
}

// This example demonstrates using channels to send and receive data
func main() {

	ch := make(chan string)
	// Generator function sends elements to the channel one by one.
	// When it is done, it closes the channel to signal the end of data.
	go func() {
		for _, str := range data {
			ch <- str
		}
		close(ch)
	}()

	// Receive strings from the channel and print them
	// For loop terminates when channel is closed
	for str := range ch {
		fmt.Println(str)
	}
}
