package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovering from panic. Type of r: %T value of r: %v\n", r, r)
		}
	}()

	panic("testing")
}
