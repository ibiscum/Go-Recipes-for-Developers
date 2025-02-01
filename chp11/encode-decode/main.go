package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Version string `json:"ver"`
	// Encoded as "Name"
	Name string
	Type string `json:"type,omitempty"` // Encoded as "type",
	// and will be omitted if empty
	Style string `json:"-"` // Not encoded
	value string // Unexported field, not encoded
	kind  string `json:"kind"` // Unexported field, not encoded

	IntValue   int     `json:"intValue,omitempty"`
	FloatValue float64 `json:"floatValue,omitempty"`

	When    *time.Time    `json:"when,omitempty"`
	HowLong time.Duration `json:"howLong,omitempty"`
}

type Enclosing struct {
	Field    string `json:"field"`
	Embedded `json:"embedded"`
}

type Embedded struct {
	Field string `json:"embeddedField"`
}

func main() {
	cfg := Config{
		Version: "1.1",
		Name:    "name",
		Type:    "example",
		Style:   "json",
		value:   "example config value",
		kind:    "test",
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	// Marshaled JSON: {"ver":"1.1","Name":"name","type":"example"}

	fmt.Println("Marshaled JSON:", string(data))
	if err := json.Unmarshal([]byte(`{
   "Ver": "1.2",
   "name": "New name",
   "value": "val"}`), &cfg); err != nil {
		panic(err)
	}
	// Unmarshaled {Version:1.2 Name:New name Type:example Style:json value:example config value
	//    kind:test When:0001-01-01 00:00:00 +0000 UTC HowLong:0s}

	fmt.Printf("Unmarshaled %+v\n", cfg)

	enc := Enclosing{
		Field: "enclosing",
		Embedded: Embedded{
			Field: "embedded",
		},
	}
	data, err = json.Marshal(enc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshaled JSON", string(data))
	// Marshaled JSON: {"field":"enclosing","embedded":{"embeddedField":"embedded"}}

	config := map[string]any{
		"ver":  "1.0",
		"Name": "config",
		"type": "example",
	}
	data, err = json.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshaled JSON", string(data))
	// Marshaled JSON {"Name":"config","type":"example","ver":"1.0"}

	numbersWithNil := []any{1, 2, nil, 3}
	data, err = json.Marshal(numbersWithNil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// [1,2,null,3]

	configurations := map[string]map[string]any{
		"cfg1": {
			"ver":  "1.0",
			"Name": "config1",
		},
		"cfg2": {
			"ver":  "1.1",
			"Name": "config2",
		},
	}
	data, err = json.Marshal(configurations)
	if err != nil {
		panic(err)
	}
	// {"cfg1":{"Name":"config1","ver":"1.0"},"cfg2":{"Name":"config2","ver":"1.1"}}
	fmt.Println(string(data))

	// Decode using json.Number
	decoder := json.NewDecoder(strings.NewReader(`[1.1,2,3,4.4]`))
	decoder.UseNumber()
	var output interface{}
	err = decoder.Decode(&output)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
	// [1.1 2 3 4.4]
}
