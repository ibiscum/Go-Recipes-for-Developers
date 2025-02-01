package main

import (
	"crypto/tls"
	"flag"
	"io"
	"net/http"
)

var (
	address     = flag.String("a", ":4433", "Address to listen")
	certificate = flag.String("c", "../server.crt", "Certificate file")
	key         = flag.String("k", "../privatekey.pem", "Private key")
)

type echoHandler struct{}

func (echoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.Copy(w, req.Body)
}

func main() {
	flag.Parse()
	handler := echoHandler{}

	cert, err := tls.LoadX509KeyPair(*certificate, *key)
	if err != nil {
		panic(err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	server := http.Server{
		Addr:      *address,
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	server.ListenAndServeTLS("", "")
}
