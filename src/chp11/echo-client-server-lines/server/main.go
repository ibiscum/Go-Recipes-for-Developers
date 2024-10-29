package main

import (
	"bufio"
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

// Limit line length to 1K.
const MaxLineLength = 1024

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Wrap the connection with a limited reader
	// to prevent the client from sending unbounded
	// amount of data
	limiter := &io.LimitedReader{
		R: conn,
		N: MaxLineLength + 1, // Read one extra byte to detect long lines
	}
	reader := bufio.NewReader(limiter)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				// Some error other than end-of-stream
				fmt.Println(err)
				return
			}
			// End of stream. It could be because the line is too long
			if limiter.N == 0 {
				// Line was too long
				fmt.Println("Received a line that is too long")
				return
			}
			// End of stream
			return
		}
		// Reset the limiter, so the next line can be read with
		// newlimit
		limiter.N = MaxLineLength + 1
		// Process the line
		if _, err := conn.Write(bytes); err != nil {
			fmt.Println(err)
			return
		}
	}
}
