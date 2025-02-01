package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
)

var (
	tlsAddress      = flag.String("a", ":4433", "TLS address to listen")
	serverAddresses = flag.String("s", ":8080", "Server addresses, comma separated")
	certificate     = flag.String("c", "../server.crt", "Certificate file")
	key             = flag.String("k", "../privatekey.pem", "Private key")
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
	// Create the tls listener
	tlsListener, err := tls.Listen("tcp", *tlsAddress, config)
	if err != nil {
		panic(err)
	}
	defer tlsListener.Close()
	fmt.Println("Listening TLS on ", tlsListener.Addr())

	// Listen to incoming TLS connections
	servers := strings.Split(*serverAddresses, ",")
	fmt.Println("Forwarding to servers: ", servers)

	// Listen to incoming TLS connections
	nextServer := 0
	for {
		conn, err := tlsListener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		retries := 0
		for {
			// Select the next server
			server := servers[nextServer]
			nextServer++
			if nextServer >= len(servers) {
				nextServer = 0
			}
			// Start a connection to this server
			targetConn, err := net.Dial("tcp", server)
			if err != nil {
				retries++
				fmt.Errorf("Cannot connect to %s", server)
				if retries > len(servers) {
					panic("None of the servers are available")
				}
				continue
			}
			// Start the proxy
			go handleProxy(conn, targetConn)
		}
	}
}

func handleProxy(conn, targetConn net.Conn) {
	defer conn.Close()
	defer targetConn.Close()
	// Copy data from the client to the server
	go io.Copy(targetConn, conn)
	// Copy data from the server to the client
	io.Copy(conn, targetConn)
}
