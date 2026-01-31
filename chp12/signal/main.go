package main

import (
	"context"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a simple HTTP echo service
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		escaped := html.EscapeString(string(body))
		io.WriteString(w, escaped)
	})
	server := &http.Server{Addr: ":8080"}

	// Listen for SIGINT and SIGTERM signals
	// Terminate the server with the signal
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigTerm
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.
			Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	server.ListenAndServe()
}
