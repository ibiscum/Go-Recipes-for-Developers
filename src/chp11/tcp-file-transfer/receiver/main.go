package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

var address = flag.String("a", ":8008", "Address to listen")

type fileMetadata struct {
	Size    uint64
	Mode    uint32
	NameLen uint16
}

func main() {
	flag.Parse()
	os.Mkdir("downloads", 0770)
	// Create a TCP listener
	listener, err := net.Listen("tcp", *address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on ", listener.Addr())
	// Listen to incoming TCP connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Read the file metadata
	var meta fileMetadata
	err := binary.Read(conn, binary.LittleEndian, &meta)
	if err != nil {
		fmt.Println(err)
		return
	}
	if meta.NameLen > 255 {
		fmt.Println("File name too long")
		return
	}
	// Read the file name
	name := make([]byte, meta.NameLen)
	_, err = io.ReadFull(conn, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create the file
	file, err := os.OpenFile(
		filepath.Join("downloads", string(name)),
		os.O_CREATE|os.O_WRONLY,
		os.FileMode(meta.Mode),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// Copy the file contents
	_, err = io.CopyN(file, conn, int64(meta.Size))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received file %s: %d bytes\n", string(name), meta.Size)
}
