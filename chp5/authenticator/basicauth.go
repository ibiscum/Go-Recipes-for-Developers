package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// BasicAuthCredentials are HTTP basic auth username:password
// credentials
type BasicAuthCredentials struct {
	Name     string
	Password string
}

// Serialize returns []byte(base64(username:password))
func (cred BasicAuthCredentials) Serialize() []byte {
	str := fmt.Sprintf("%s:%s", cred.Name, cred.Password)
	target := make([]byte, base64.RawURLEncoding.EncodedLen(len(str)))
	base64.RawURLEncoding.Encode(target, []byte(str))
	return target
}

// Type returns "basicauth"
func (BasicAuthCredentials) Type() string { return "basicauth" }

type BasicAuthParser struct{}

// Parse parses a basic auth credential obtained by Serialize method
func (BasicAuthParser) Parse(input []byte) (Credentials, error) {
	target := make([]byte, base64.RawURLEncoding.DecodedLen(len(input)))
	_, err := base64.RawURLEncoding.Decode(target, input)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(target), ":")
	if len(parts) != 2 {
		return nil, ErrBadCredentials
	}
	return BasicAuthCredentials{
		Name:     parts[0],
		Password: parts[1],
	}, nil
}

// BasicAuthAuthenticator authenticates a user based on a map
type BasicAuthAuthenticator struct {
	// Map of known user-passwords
	users map[string]string
}

func (auth BasicAuthAuthenticator) Login(credentials Credentials) (Session, error) {
	basicAuthCredentials, ok := credentials.(BasicAuthCredentials)
	if !ok {
		return nil, ErrUnauthorized
	}
	pwd, ok := auth.users[basicAuthCredentials.Name]
	if !ok {
		return nil, ErrUnauthorized
	}
	if pwd != basicAuthCredentials.Password {
		return nil, ErrUnauthorized
	}
	return &BasicAuthSession{
		User: basicAuthCredentials.Name,
		open: true,
	}, nil
}

// BasicAuthSession represents an active user session
type BasicAuthSession struct {
	User string

	open bool
}

// UserID returns the user id of the session
func (session BasicAuthSession) UserID() string {
	return session.User
}

// Close closes the session
func (session BasicAuthSession) Close() {
	// session.open = false
}

type BasicAuthFactory struct{}

func (BasicAuthFactory) NewInstance() Authenticator {
	return &BasicAuthAuthenticator{
		users: map[string]string{
			"john": "doe",
			"jane": "doe",
			"foo":  "bar",
		},
	}
}

func init() {
	RegisterAuthenticator("basic", BasicAuthFactory{})
}
