package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

const downloadDir = "downloads"

// This example demostrates managing multiple independent
// goroutines. A wait group is used to wait for all goroutines to
// end. Each goroutine downloads a file and stores it in the file
// system. There are no dependencies or communications between
// goroutines.
func main() {
	urls := []string{
		"https://pkg.go.dev/bufio",
		"https://pkg.go.dev/builtin",
		"https://pkg.go.dev/bytes",
		"https://pkg.go.dev/cmp",
		"https://pkg.go.dev/context",
		"https://pkg.go.dev/crypto",
		"https://pkg.go.dev/embed",
		"https://pkg.go.dev/encoding",
		"https://pkg.go.dev/errors",
		"https://pkg.go.dev/expvar",
		"https://pkg.go.dev/flag",
		"https://pkg.go.dev/hash",
		"https://pkg.go.dev/log",
		"https://pkg.go.dev/plugin",
	}

	wg := sync.WaitGroup{}

	for _, u := range urls {
		wg.Add(1)
		go func(downloadURL string) {
			defer wg.Done()

			rsp, err := http.Get(downloadURL)
			if err != nil {
				log.Printf("Cannot download %s: %s", downloadURL, err)
				return
			}
			defer rsp.Body.Close()
			if rsp.StatusCode != 200 {
				log.Printf("Cannot download %s: %s", downloadURL, rsp.Status)
				return
			}
			fname := path.Base(rsp.Request.URL.Path)
			// Ensure the derived filename does not contain path separators or parent directory references.
			if strings.Contains(fname, "/") || strings.Contains(fname, "\\") || strings.Contains(fname, "..") || fname == "" {
				log.Printf("Refusing to use unsafe filename %q for URL %s", fname, downloadURL)
				return
			}
			// Ensure the download directory exists and write files into it.
			if err := os.MkdirAll(downloadDir, 0o755); err != nil {
				log.Printf("Cannot create download directory %s: %s", downloadDir, err)
				return
			}
			fullPath := filepath.Join(downloadDir, fname)
			file, err := os.Create(fullPath)
			if err != nil {
				log.Printf("Cannot write file %s: %s", fullPath, err)
				return
			}
			defer file.Close()
			if _, err := io.Copy(file, rsp.Body); err != nil {
				log.Printf("Cannot read %s: %s", downloadURL, err)
				return
			}
		}(u)
	}
	// Wait for all goroutines to end
	wg.Wait()
}
