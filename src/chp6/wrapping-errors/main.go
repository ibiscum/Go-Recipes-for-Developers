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

func LoadConfig(f string) (*Config, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("file %s: %w", f, err)
	}
	defer file.Close()
	var cfg Config
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("While unmarshaling %s: %w", f, err)
	}
	return &cfg, nil
}

func main() {
	_, err := LoadConfig("nonexistant-config.json")
	fmt.Printf("Error: %v, Is ErrNotExist: %v\n", err, errors.Is(err, os.ErrNotExist))
	_, err = LoadConfig("invalid-config.json")
	fmt.Printf("Error: %v\n", err)
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		fmt.Printf("Unwrapped syntax error: %+v\n", syntaxError)
	}
	config, err := LoadConfig("valid-config.json")
	fmt.Println(config, err)
}
