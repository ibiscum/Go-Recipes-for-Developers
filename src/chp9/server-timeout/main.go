package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

const RequestTimeout = time.Second * 5

func handleRequest(ctx context.Context, c net.Conn) {
	defer c.Close()
	for {
		if ctx.Err() != nil {
			fmt.Println("Canceled")
			return
		}
		fmt.Println("Waiting")
		time.Sleep(time.Second)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on localhost:8080. Run curl localhost:8080 to establish connection...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go func(c net.Conn) {
			// Step 1:
			// Request times out after duration: RequestTimeout
			ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)

			// Step 2:
			// Make sure cancel is called
			defer cancel()

			// Step 3:
			// Pass the context to handler
			handleRequest(ctx, c)
		}(conn)
	}
}
