package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

type Data struct {
	IntValue   int64
	BoolValue  bool
	ArrayValue [2]int64
}

func EncodeString(s string) []byte {
	// Allocate the output buffer for string length (int16) + len(string)
	buffer := make([]byte, len(s)+2)
	// Encode the length little endian - 2 bytes
	binary.LittleEndian.PutUint16(buffer, uint16(len(s)))
	// Copy the string bytes
	copy(buffer[2:], []byte(s))
	return buffer
}

func DecodeString(input []byte) (string, error) {
	// Read the string length. It must be at least 2 bytes
	if len(input) < 2 {
		return "", fmt.Errorf("invalid input")
	}
	n := binary.LittleEndian.Uint16(input)
	if int(n)+2 > len(input) {
		return "", fmt.Errorf("invalid input")
	}
	return string(input[2 : n+2]), nil
}

func main() {
	output := bytes.Buffer{}
	data := Data{
		IntValue:   1,
		BoolValue:  true,
		ArrayValue: [2]int64{1, 2},
	}
	err := binary.Write(&output, binary.BigEndian, data)
	if err != nil {
		log.Fatal(err)
	}
	stream := output.Bytes()
	fmt.Printf("Big endian encoded data   : %v\n", stream)
	var value1 Data
	err = binary.Read(bytes.NewReader(stream), binary.BigEndian, &value1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded data: %v\n", value1)

	output = bytes.Buffer{}
	err = binary.Write(&output, binary.LittleEndian, data)
	if err != nil {
		log.Fatal(err)
	}
	stream = output.Bytes()
	fmt.Printf("Little endian encoded data: %v\n", stream)
	var value2 Data
	err = binary.Read(bytes.NewReader(stream), binary.LittleEndian, &value2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded data: %v\n", value2)

	encodedString := EncodeString("Hello world!")
	fmt.Printf("Encoded string: %v\n", encodedString)
	decodedString, _ := DecodeString(encodedString)
	fmt.Printf("Decoded string: %s\n", decodedString)
}
