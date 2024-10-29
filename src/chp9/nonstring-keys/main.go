package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Key is an uint that is encoded as an hex strings for JSON key
type Key uint

func (k *Key) UnmarshalText(data []byte) error {
	v, err := strconv.ParseInt(string(data), 16, 64)
	if err != nil {
		return err
	}
	*k = Key(v)
	return nil
}

func (k Key) MarshalText() ([]byte, error) {
	s := strconv.FormatUint(uint64(k), 16)
	return []byte(s), nil
}

func main() {
	input := `{
    "13AD": "5037",
    "3E22": "15906",
    "90A3": "37027"
  }`

	var data map[Key]string
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	d, err := json.Marshal(map[Key]any{
		Key(123): "123",
		Key(255): "255",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}
