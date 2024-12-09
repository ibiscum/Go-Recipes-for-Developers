package main

import (
	"encoding/json"
	"fmt"
)

type Key int64

func main() {
	var m map[Key]int
	err := json.Unmarshal([]byte(`{"123":123}`), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m[123]) // Prints 123
}
