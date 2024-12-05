package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	address     = flag.String("a", "https://localhost:4433", "Server address")
	certificate = flag.String("c", "../server.crt", "Certificate file")
)

func main() {
	flag.Parse()
	certData, err := os.ReadFile(*certificate)
	if err != nil {
		panic(err)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(certData)
	if !ok {
		panic("failed to parse root certificate")
	}
	config := tls.Config{
		RootCAs: roots,
	}
	transport := &http.Transport{
		TLSClientConfig: &config,
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Post(*address, "text/plain", strings.NewReader("ping\n"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
