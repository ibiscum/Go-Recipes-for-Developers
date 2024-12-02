package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidTypeAndID = errors.New("invalid TypeAndID")

// TypeAndID is encoded to JSON as type:id
type TypeAndID struct {
	Type string
	ID   int
}

// Implementation of json.Marshaler
func (t TypeAndID) MarshalJSON() (out []byte, err error) {
	s := fmt.Sprintf(`"%s:%d"`, t.Type, t.ID)
	out = []byte(s)
	return
}

// Implementation of json.Unmarshaler. Note the pointer receiver
func (t *TypeAndID) UnmarshalJSON(in []byte) (err error) {
	if in[0] != '"' || in[len(in)-1] != '"' {
		err = ErrInvalidTypeAndID
		return
	}
	in = in[1 : len(in)-1]
	parts := strings.Split(string(in), ":")
	if len(parts) != 2 {
		err = ErrInvalidTypeAndID
		return
	}
	// The second part must be a valid integer
	t.ID, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}
	t.Type = parts[0]
	return
}

func main() {
	data, err := json.Marshal(TypeAndID{
		Type: "test1",
		ID:   1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	var tid TypeAndID
	if err := json.Unmarshal([]byte(`"test2:200"`), &tid); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tid)
}
