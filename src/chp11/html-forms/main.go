package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

const form = `
<!doctype html>
<html lang="en">
  <head>
  </head>
<body>
{{if .error}}Error: {{.error}}{{end}}
<form method="POST" action="/auth/login">
<input type="text" name="userName">
<input type="password" name="password">
<button type="submit">Submit</button>
</form>
</body>
</html>
`

var loginFormTemplate = template.Must(template.New("login").Parse(form))

type Authenticator struct{}

func (Authenticator) Authenticate(name, password string) (*http.Cookie, error) {
	if name == "test" && password == "password" {
		return &http.Cookie{
			Name:   "test",
			Value:  name,
			MaxAge: 3600,
			Path:   "/",
		}, nil
	}
	return nil, errors.New("Unknown user")
}

type UserHandler struct {
	Auth Authenticator
}

func (h UserHandler) HandleLogin(w http.ResponseWriter, req *http.Request) {
	// Parse the submitted form. This fills up req.PostForm
	// with the submitted information
	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the submitted fields
	userName := req.PostForm.Get("userName")
	password := req.PostForm.Get("password")
	// Handle the login request, and get a cookie
	cookie, err := h.Auth.Authenticate(userName, password)
	if err != nil {
		// Send the user back to login page, setting an error
		// cookie containing an error message
		http.SetCookie(w, h.NewErrorCookie("Username or password invalid"))
		http.Redirect(w, req, "/login.html", http.StatusFound)
		return
	}
	// Set the cookie representing user session
	http.SetCookie(w, cookie)
	// Redirect the user to the main page
	http.Redirect(w, req, "/dashboard.html", http.StatusFound)
}

func (h UserHandler) ShowLoginPage(w http.ResponseWriter, req *http.Request) {
	loginFormData := map[string]any{}
	cookie, err := req.Cookie("error_cookie")
	fmt.Println(cookie, err)
	if err == nil {
		loginFormData["error"] = cookie.Value
		// Unset the cookie
		http.SetCookie(w, &http.Cookie{
			Name:   "error_cookie",
			MaxAge: 0,
		})
	}
	w.Header().Set("Content-Type", "text/html")
	loginFormTemplate.Execute(w, loginFormData)
}

func (h UserHandler) ShowDashboardPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("success"))
}

func (h UserHandler) NewErrorCookie(msg string) *http.Cookie {
	return &http.Cookie{
		Name:   "error_cookie",
		Value:  msg,
		MaxAge: 60, // Cookie lives for 60 seconds
		Path:   "/",
	}
}

func main() {
	authenticator := Authenticator{}
	userHandler := UserHandler{
		Auth: authenticator,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/login", userHandler.HandleLogin)
	mux.HandleFunc("GET /login.html", userHandler.ShowLoginPage)
	mux.HandleFunc("GET /dashboard.html", userHandler.ShowDashboardPage)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Go to http://localhost:8080/login.html, and login as user 'test' with password 'password'")
	server.ListenAndServe()
}
