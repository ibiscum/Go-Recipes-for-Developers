package main

import (
	"log"
	"os"
)

func Logging() {
	log.Println("This is a log message similar to fmt.Println")
	log.Printf("This is a log message similar to fmt.Printf")
}

func ConfiguringPrefixes() {
	logger := log.New(log.Writer(), "prefix: ", log.LstdFlags)
	logger.Println("This is a log message with a prefix")
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Println("This is a log message with a prefix and file name")
	logger.SetFlags(log.LstdFlags | log.Llongfile)
	logger.Println("This is a log message with a prefix and long file name")
	logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	logger.Println("This is a log message with a prefix moved to the beginning of the message")
	logger.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	logger.Println("This is a log message with with UTC time")
}

func SettingOutput() {
	output, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	logger := log.New(os.Stderr, "", log.LstdFlags)
	logger.Println("This is a log message to stderr")
	logger.SetOutput(output)
	logger.Println("This is a log message to log.txt")
	logger.SetOutput(os.Stderr)
	logger.Println("Message to log.txt was written")
}

func main() {
	Logging()
	ConfiguringPrefixes()
	SettingOutput()
}
