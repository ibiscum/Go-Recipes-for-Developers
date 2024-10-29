package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Data struct {
	Value int `json:"value"`
}

func generate(nItems int) <-chan Data {
	ch := make(chan Data)
	go func() {
		for i := 0; i < nItems; i++ {
			ch <- Data{Value: i}
		}
		close(ch)
	}()
	return ch
}

func stream(out io.Writer, input <-chan Data) error {
	enc := json.NewEncoder(out)
	if _, err := out.Write([]byte{'['}); err != nil {
		return err
	}
	first := true
	for obj := range input {
		if first {
			first = false
		} else {
			if _, err := out.Write([]byte{','}); err != nil {
				return err
			}
		}
		if err := enc.Encode(obj); err != nil {
			return err
		}
	}

	if _, err := out.Write([]byte{']'}); err != nil {
		return err
	}
	return nil
}

func parse(input *json.Decoder) (output []Data, err error) {
	// Parse the array beginning delimiter
	var tok json.Token
	tok, err = input.Token()
	if err != nil {
		return
	}
	if tok != json.Delim('[') {
		err = fmt.Errorf("Array begin delimiter expected")
		return
	}
	// Parse array elements using Decode
	for {
		var data Data
		err = input.Decode(&data)
		if err != nil {
			// Decode failed. Either there is an input error, or
			// we are at the end of the stream
			tok, err = input.Token()
			if err != nil {
				// Data error
				return
			}
			// Are we at the end?
			if tok == json.Delim(']') {
				// Yes, there is no error
				err = nil
				break
			}
		}
		output = append(output, data)
	}
	return
}

func main() {
	buf := bytes.Buffer{}
	ch := generate(100)
	stream(&buf, ch)
	str := buf.String()
	fmt.Println(str)
	decoder := json.NewDecoder(strings.NewReader(str))
	result, err := parse(decoder)
	fmt.Println(result, err)
}
