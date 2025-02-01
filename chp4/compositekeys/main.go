package main

import "fmt"

type Key struct {
	UserID    string
	SessionID string
}

type User struct {
	ID   string
	Name string
}

func main() {
	compositeKeyMap := map[Key]User{}
	compositeKeyMap[Key{
		UserID:    "123",
		SessionID: "1",
	}] = User{
		Name: "John Doe",
		ID:   "123",
	}

	jane := User{
		Name: "Jane Doe",
		ID:   "124",
	}

	key := Key{
		UserID:    jane.ID,
		SessionID: "2",
	}
	compositeKeyMap[key] = jane

	for k, v := range compositeKeyMap {
		fmt.Printf("Key: %+v, Value: %+v\n", k, v)
	}
}
