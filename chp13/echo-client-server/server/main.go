package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

var address = flag.String("a", ":8008", "Address to listen")

func main() {
	flag.Parse()

	// Create a TCP listener
	listener, err := net.Listen("tcp", *address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on ", listener.Addr())
	defer listener.Close()
	// Listen to incoming TCP connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	io.Copy(conn, conn)
}
