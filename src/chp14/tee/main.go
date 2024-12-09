package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter the name of file to read")
		os.Exit(1)
	}
	// Create an http test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write data to stdout
		io.Copy(os.Stdout, r.Body)
	}))
	defer ts.Close()
	serverURL := ts.URL
	dataFile := os.Args[1]

	pipeReader, pipeWriter := io.Pipe()
	file, err := os.Open(dataFile)
	if err != nil {
		// Handle error
		panic(err)
	}
	defer file.Close()
	tee := io.TeeReader(file, pipeWriter)
	go func() {
		// Copy the file to stdout
		io.Copy(os.Stdout, pipeReader)
	}()
	_, err = http.Post(serverURL, "text/plain", tee)
	if err != nil {
		// Make sure pipe is closed
		pipeReader.Close()
	}
}
