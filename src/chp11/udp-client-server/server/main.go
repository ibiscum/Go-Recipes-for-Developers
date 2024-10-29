package main

import (
	"flag"
	"fmt"
	"net"
)

var address = flag.String("a", ":8008", "Address to listen")

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp4", *address)
	if err != nil {
		panic(err)
	}
	// Create a UDP connection
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on ", conn.LocalAddr())
	defer conn.Close()
	// Listen to incoming UDP connections
	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Received %d bytes from %s\n", n, remoteAddr)
		if n > 0 {
			_, err := conn.WriteToUDP(buf[:n], remoteAddr)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
