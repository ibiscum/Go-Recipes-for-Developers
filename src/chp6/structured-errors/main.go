package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Option string `json:"option"`
}

type ErrFile struct {
	Name string
	When string
	Err  error
}

var cfg Config

func (err ErrFile) Error() string {
	return fmt.Sprintf("%s: file %s, when %s", err.Err, err.Name, err.When)
}

func (err ErrFile) Is(e error) bool {
	_, ok := e.(ErrFile)
	return ok
}

func (err ErrFile) Unwrap() error { return err.Err }

func ReadConfigFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return ErrFile{
			Name: name,
			Err:  err,
			When: "opening configuration file",
		}
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return fmt.Errorf("While unmarshaling %s: %w", name, err)
	}
	return nil
}

func main() {
	err := ReadConfigFile("nonexistant-config.json")
	fmt.Printf("Error: %v\n", err)
	var errFile ErrFile
	if errors.As(err, &errFile) {
		fmt.Printf("Errfile: %v\n", errFile)
	}
	err = ReadConfigFile("invalid-config.json")
	fmt.Printf("Error: %v\n", err)
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		fmt.Printf("Unwrapped syntax error: %+v\n", syntaxError)
	}
	err = ReadConfigFile("valid-config.json")
	fmt.Println(cfg, err)
}
