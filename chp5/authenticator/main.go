package main

import (
	"errors"
	"fmt"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrBadCredentials = errors.New("bad credentials")

// Authenticator uses implementation-specific credentials to create an
// implementation-specific session
type Authenticator interface {
	Login(credentials Credentials) (Session, error)
}

// Credentials contains the credentials to authenticate a user to the backend
type Credentials interface {
	Serialize() []byte
	Type() string
}

// CredentialParse implementation parses backend-specific credentials from []byte input
type CredentialParser interface {
	Parse([]byte) (Credentials, error)
}

// A backend-specific session identifies the user and provides a way to close the session
type Session interface {
	UserID() string
	Close()
}

func main() {
	// Create a new authenticator
	authenticator := NewInstance("basic")

	session, err := authenticator.Login(BasicAuthCredentials{
		Name:     "john",
		Password: "doe",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Login john/doe, session: %+v\n", session)

	_, err = authenticator.Login(BasicAuthCredentials{
		Name:     "john",
		Password: "password",
	})
	if err != nil {
		fmt.Printf("Login john/password: err: %v\n", err)
	}

}
