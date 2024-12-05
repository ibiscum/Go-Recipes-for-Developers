package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

func usingDefaultClient() error {
	response, err := http.Get("http://example.com")
	if err != nil {
		// Handle error
		return err
	}
	// Always close response body
	defer response.Body.Close()
	if response.StatusCode/100 == 2 {
		// HTTP 2xx, call was successful.
		// Work with response.Body
		io.Copy(os.Stdout, response.Body)
	}
	return nil
}

func usingCustomClient() error {
	client := http.Client{
		// Set a timeout for all outgoing calls.
		// If the call does not complete within 30 seconds, timeout.
		Timeout: 30 * time.Second,
	}
	response, err := client.Get("http://example.com")
	if err != nil {
		// handle error
		return err
	}
	// Always close response body
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
	return nil
}

func main() {
	if err := usingDefaultClient(); err != nil {
		panic(err)
	}
	if err := usingCustomClient(); err != nil {
		panic(err)
	}
}
