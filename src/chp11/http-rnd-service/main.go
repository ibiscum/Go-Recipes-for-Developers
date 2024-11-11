package main

import (
	"encoding/binary"
	"io"
	"net/http"
	"os"
	"strconv"
)

type RandomService struct {
	rndSource io.Reader
}

func (svc RandomService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Read 4 bytes from the random number source, convert it to string
	data := make([]byte, 4)
	_, err := svc.rndSource.Read(data)
	if err != nil {
		// This will return an HTTP 500 error with the error message
		// as the message body
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Decode random data using binary little endian encoding
	value := binary.LittleEndian.Uint32(data)
	// Write the data to the output
	w.Write([]byte(strconv.Itoa(int(value))))
}

func main() {
	file, err := os.Open("/dev/random")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	svc := RandomService{
		rndSource: file,
	}
	mux := http.NewServeMux()
	mux.Handle("GET /rnd", svc)
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
