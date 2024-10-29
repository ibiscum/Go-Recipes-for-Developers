package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var (
	address = flag.String("address", ":8080", "HTTP server address")
	dir     = flag.String("dir", "/usr/share/doc", "Directory to serve")
	ext     = flag.String("ext", "", "File extension to serve")
)

type filterFS struct {
	dirFS http.FileSystem
	ext   string
}

func (f filterFS) Open(name string) (http.File, error) {
	if filepath.Ext(name) != f.ext {
		return nil, os.ErrNotExist
	}
	return f.dirFS.Open(name)
}

func main() {
	flag.Parse()
	server := http.Server{
		Addr: *address,
	}
	var fs http.FileSystem
	if len(*ext) > 0 {
		fs = filterFS{
			dirFS: http.Dir(*dir),
			ext:   *ext,
		}
	} else {
		fs = http.Dir(*dir)
	}
	server.Handler = http.FileServer(fs)
	fmt.Printf("Serving %s on HTTP %s\n", *dir, *address)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
