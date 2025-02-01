package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const form = `
<!doctype html>
<html lang="en">
  <head>
  </head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
<input type="text" name="textField">
<input type="file" name="fileField">
<button type="submit">submit</button>
</form>
</body>
</html>
`

func handleUpload(w http.ResponseWriter, request *http.Request) {
	reader, err := request.MultipartReader()
	if err != nil {
		http.Error(w, "Not a multipart request", http.StatusBadRequest)
		return
	}
	for {
		part, err := reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		formValues := make(url.Values)
		if fileName := part.FileName(); fileName != "" {
			// This part contains a file
			output, err := os.Create(filepath.Join("uploads", fileName))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer output.Close()
			fmt.Println("Uploading file", fileName)
			if _, err := io.Copy(output, part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if fieldName := part.FormName(); fieldName != "" {
			// This part contains form data for an input field
			data, err := io.ReadAll(part)
			if err != nil {
				// Handle error
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			formValues[fieldName] = append(formValues[fieldName],
				string(data))
		}
	}
	http.Redirect(w, request, "/upload.html", http.StatusFound)
}

func main() {
	os.Mkdir("uploads", 0o775)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /upload", handleUpload)
	mux.HandleFunc("GET /upload.html", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(form))
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	fmt.Println("Go to http://localhost:8081/upload.html")
	server.ListenAndServe()
}
