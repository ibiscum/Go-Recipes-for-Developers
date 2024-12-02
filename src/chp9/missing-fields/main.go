package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type APIRequest struct {
	// If type is not specified, it will be nil
	Type *string `json:"type"`
	// There will be a default value for seq
	Seq int `json:"seq"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	req := APIRequest{
		Seq: 1, // Set the default value
	}
	if err := json.Unmarshal(data, &req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Check which fields are provided
	if req.Type != nil {
		fmt.Println("Type specified", *req.Type)
	} else {
		fmt.Println("No type was specified")
	}

	// If seq is provided in the input, req.Seq will be set to that value. Otherwise, it will be 1.
	if req.Seq == 1 {
		fmt.Println("Sequence", 1)
	} else {
		fmt.Println("Sequence", req.Seq)
	}
}

func main() {
	// Simulate request handling: empty body
	req, err := http.NewRequest(http.MethodPost, "localhost", strings.NewReader(`{}`))
	if err != nil {
		panic(err)
	}
	handler(nil, req)

	// Body with type specified
	req, err = http.NewRequest(http.MethodPost, "localhost", strings.NewReader(`{"type":"test1"}`))
	if err != nil {
		panic(err)
	}
	handler(nil, req)

	// Body with type and sequence specified
	req, err = http.NewRequest(http.MethodPost, "localhost", strings.NewReader(`{"type":"test2", "seq":100}`))
	if err != nil {
		panic(err)
	}
	handler(nil, req)

}
