package main

type AuthenticatorFactory interface {
	NewInstance() Authenticator
}

var registry = map[string]AuthenticatorFactory{}

// RegisterAuthenticator registers a new authenticator factory
func RegisterAuthenticator(name string, factory AuthenticatorFactory) {
	registry[name] = factory
}

func NewInstance(authType string) Authenticator {
	// Create a new instance using the selected factory.
	// If the given authType has not been registered, this will panic
	return registry[authType].NewInstance()
}
