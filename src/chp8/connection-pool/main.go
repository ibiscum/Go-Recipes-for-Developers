package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func server() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

type ConnectionPool struct {
	available chan net.Conn
	total     chan struct{}
}

func (pool *ConnectionPool) GetConnection() (net.Conn, error) {
	select {
	// If there are connections available in the pool, return one
	case conn := <-pool.available:
		fmt.Printf("Returning an idle connection.\n")
		return conn, nil

	default:
		// No connections are available
		select {
		case conn := <-pool.available:
			fmt.Printf("Returning an idle connection.\n")
			return conn, nil

		case pool.total <- struct{}{}: // Wait until pool is not full
			fmt.Println("Creating a new connection")
			// Create a new connection
			conn, err := net.Dial("tcp", "localhost:2000")
			if err != nil {
				return nil, err
			}
			return conn, nil
		}
	}
}

func (pool *ConnectionPool) Release(conn net.Conn) {
	pool.available <- conn
	fmt.Printf("Releasing a connection. \n")
}

func (pool *ConnectionPool) Close(conn net.Conn) {
	fmt.Println("Closing connection")
	conn.Close()
	<-pool.total
}

const PoolSize = 10

func main() {
	go server()

	pool := ConnectionPool{
		available: make(chan net.Conn, PoolSize),
		total:     make(chan struct{}, PoolSize),
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Get a new connection from the pool
			conn, err := pool.GetConnection()
			if err != nil {
				log.Fatal(err)
			}
			// Work with the connection
			_, err = conn.Write([]byte("test"))
			if err != nil {
				pool.Close(conn)
				return
			}
			buf := make([]byte, 4)
			_, err = io.ReadFull(conn, buf)
			if err != nil {
				pool.Close(conn)
				return
			}
			pool.Release(conn)
		}(i)
	}
	wg.Wait()
}
