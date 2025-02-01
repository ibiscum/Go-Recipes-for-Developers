package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Users struct {
	Users map[string]User `json:"users"`
}

func main() {
	input := `{
  "users": {
      "abb64dfe-d4a8-47a5-b7b0-7613fe3fd11f": {
         "name": "John",
         "type": "admin"
      },
      "b158161c-0588-4c67-8e4b-c07a8978f711": {
         "name": "Amy",
         "type": "editor"
      }
   }
  }`
	var users Users
	if err := json.Unmarshal([]byte(input), &users); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)
}
