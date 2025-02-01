package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type requestIDKeyType int

var requestIDKey requestIDKeyType

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func GetRequestID(ctx context.Context) string {
	id, _ := ctx.Value(requestIDKey).(string)
	return id
}

type Privilege uint16

const (
	PrivilegeRead  Privilege = 0x0001
	PrivilegeWrite Privilege = 0x0002
)

func GetPrivileges(userID string) map[string]Privilege {
	// Return mock privileges for the user
	return map[string]Privilege{
		"object": PrivilegeRead | PrivilegeWrite,
	}
}

type authInfoKeyType int

var authInfoKey authInfoKeyType

type AuthInfo struct {
	sync.Mutex
	UserID     string
	privileges map[string]Privilege
}

func GetAuthInfo(ctx context.Context) *AuthInfo {
	info, _ := ctx.Value(authInfoKey).(*AuthInfo)
	return info
}

func (auth *AuthInfo) GetPrivileges() map[string]Privilege {
	// Use mutex to initialize the privileges in a thread-safe way
	auth.Lock()
	defer auth.Unlock()
	if auth.privileges == nil {
		auth.privileges = GetPrivileges(auth.UserID)
	}
	return auth.privileges
}

// Mock authenticator function
func authenticate(_ *http.Request) (*AuthInfo, error) {
	return &AuthInfo{
		UserID:     uuid.New().String(),
		privileges: map[string]Privilege{},
	}, nil
}

// Authentication middleware
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Authenticate the caller
			var authInfo *AuthInfo
			var err error
			authInfo, err = authenticate(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			authInfo.privileges = GetPrivileges(authInfo.UserID)
			// Create a new context with the authentication info
			newCtx := context.WithValue(r.Context(), authInfoKey, authInfo)
			// Pass the new context to the next handler
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func main() {
	// Create a new context with a request id
	ctx := WithRequestID(context.Background(), uuid.New().String())

	fmt.Printf("Request id in ontext.Background: %s\n", GetRequestID(context.Background()))
	fmt.Printf("Request id in context: %s\n", GetRequestID(ctx))

	// Simulate HTTP call here:

	mw := AuthMiddleware()

	req, err := http.NewRequest(http.MethodGet, "localhost:00", nil)
	if err != nil {
		panic(err)
	}
	mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authInfo := GetAuthInfo(r.Context())
		fmt.Printf("AuthInfo in context: %+v\n", authInfo)
	})).ServeHTTP(nil, req)

}
