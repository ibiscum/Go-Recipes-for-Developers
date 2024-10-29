package main

import (
	"encoding/binary"
	"flag"
	"io"
	"net"
	"os"
)

var address = flag.String("a", ":8008", "Server address")
var file = flag.String("file", "", "File to send")

type fileMetadata struct {
	Size    uint64
	Mode    uint32
	NameLen uint16
}

func main() {
	flag.Parse()

	file, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		panic(err)
	}

	// Encode file metadata
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	md := fileMetadata{
		Size:    uint64(fileInfo.Size()),
		Mode:    uint32(fileInfo.Mode()),
		NameLen: uint16(len(fileInfo.Name())),
	}
	if err := binary.Write(conn, binary.LittleEndian, md); err != nil {
		panic(err)
	}
	// The file name
	if _, err := conn.Write([]byte(fileInfo.Name())); err != nil {
		panic(err)
	}
	// The file contents
	if _, err := io.Copy(conn, file); err != nil {
		panic(err)
	}
	conn.Close()
}
