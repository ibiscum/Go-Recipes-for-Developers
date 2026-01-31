package main

import (
	"crypto/tls"
	"flag"
	"io"
	"net/http"
	"html"
)

var (
	address     = flag.String("a", ":4433", "Address to listen")
	certificate = flag.String("c", "../server.crt", "Certificate file")
	key         = flag.String("k", "../privatekey.pem", "Private key")
)

type echoHandler struct{}

func (echoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain") // Explicitly set content type
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	escapedBody := html.EscapeString(string(body)) // Escape user-provided input
	w.Write([]byte(escapedBody))
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
