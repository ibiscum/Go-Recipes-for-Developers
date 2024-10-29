package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
)

var (
	address     = flag.String("a", ":4433", "Address to listen")
	certificate = flag.String("c", "../server.crt", "Certificate file")
	key         = flag.String("k", "../privatekey.pem", "Private key")
)

func main() {
	flag.Parse()

	// Load the key pair
	cer, err := tls.LoadX509KeyPair(*certificate, *key)
	if err != nil {
		panic(err)
	}
	// Create TLS configuration for the listener
	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
	}
	// Create the listener
	listener, err := tls.Listen("tcp", *address, config)
	if err != nil {
		panic(err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening TLS on ", listener.Addr())
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
