package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"os"
)

var (
	addr     = flag.String("addr", "", "Server address")
	certFile = flag.String("cert", "../server.crt", "TLS certificate file")
)

func main() {
	flag.Parse()

	// Create new certificate pool
	roots := x509.NewCertPool()
	// Load server certificate
	certData, err := os.ReadFile(*certFile)
	if err != nil {
		panic(err)
	}
	ok := roots.AppendCertsFromPEM(certData)
	if !ok {
		panic("failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", *addr, &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		panic(err)
	}
	// Send a line of text
	text := []byte("Hello echo server!")
	conn.Write(text)
	// Read the response
	response := make([]byte, len(text))
	conn.Read(response)
	fmt.Println(string(response))
	conn.Close()
}
