package main

import (
	"flag"
	"fmt"
	"net"
)

var address = flag.String("a", ":8008", "Server address")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		panic(err)
	}
	// Send a line of text
	text := []byte("Hello echo server!\n")
	conn.Write(text)
	// Read the response
	response := make([]byte, len(text))
	conn.Read(response)
	fmt.Println(string(response))
	// Launch an attack by sending large amount of data
	data := make([]byte, 1024*1000)
	_, err = conn.Write(data)
	fmt.Println(err)

	conn.Close()
}
