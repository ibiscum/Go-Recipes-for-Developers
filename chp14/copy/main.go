package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	sourceFile = flag.String("src", "", "Source file")
	targetFile = flag.String("tgt", "", "Target file")
)

func help() {
	fmt.Println("Run with -src <sourceFile> -tgt <targetFile>")
	os.Exit(1)
}

func main() {
	flag.Parse()

	sourceFileName := *sourceFile
	if sourceFileName == "" {
		help()
	}
	targetFileName := *targetFile
	if targetFileName == "" {
		help()
	}
	if sourceFileName == targetFileName {
		panic("Target is the same as source")
	}
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, sourceFile); err != nil {
		panic(err)
	}
}
