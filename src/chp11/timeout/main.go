package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Create a TCP listener
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on ", listener.Addr())
	defer listener.Close()

	// Start server
	go func() {
		// Listen to incoming TCP connections
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			go handleConnectionWithTimeout(conn)
		}
	}()

	// Connect to server
	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		panic(err)
	}

	// Write data
	clientConn.Write([]byte("Hello server!"))
	data := make([]byte, 1024)
	n, err := clientConn.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	// Trigger timeout
	time.Sleep(3 * time.Second)

	clientConn.Write([]byte("Hello server after timeout!"))
	n, err = clientConn.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))

}

func handleConnectionWithTimeout(conn net.Conn) {
	for {
		// Set a 1 second read timeout
		conn.SetReadDeadline(time.Now().Add(time.Second))
		_, err := io.Copy(conn, conn)
		if err != nil {
			if errors.Is(err, os.ErrDeadlineExceeded) {
				fmt.Println("Read timeout, restarting")
			} else {
				return
			}
		}
	}
}
