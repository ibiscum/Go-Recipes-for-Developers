package main

import (
	"flag"
	"fmt"
	"net"
)

var address = flag.String("a", "localhost:8008", "Server address")

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp4", *address)
	if err != nil {
		panic(err)
	}
	// Create a UDP connection, local address chosen randomly
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("UDP server %s\n", conn.RemoteAddr())
	defer conn.Close()
	// Send a line of text
	text := []byte("Hello echo server!")
	n, err := conn.Write(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Written %d bytes\n", n)
	// Read the response
	response := make([]byte, 1024)
	conn.ReadFromUDP(response)
	fmt.Println(string(response))
	conn.Close()
}
