package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

func main() {
	// Create an http test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write data to stdout
		io.Copy(os.Stdout, r.Body)
	}))
	defer ts.Close()
	serverURL := ts.URL
	// Create a test payload
	payload := map[string]any{
		"key": "value",
	}

	pipeReader, pipeWriter := io.Pipe()
	go func() {
		// Close the writer side, so the reader knows when it is done
		defer pipeWriter.Close()
		encoder := json.NewEncoder(pipeWriter)
		if err := encoder.Encode(payload); err != nil {
			if errors.Is(err, io.ErrClosedPipe) {
				// The reader side terminated with error
				panic(err)
			} else {
				// Handle error
				panic(err)
			}
		}
	}()
	if _, err := http.Post(serverURL, "application/json", pipeReader); err != nil {
		// Close the reader, so the writing goroutine terminates
		pipeReader.Close()
		// Handle error
	}
}
