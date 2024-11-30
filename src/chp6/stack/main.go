package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			stackTrace := string(debug.Stack())
			// Work with stackTrace
			fmt.Println(stackTrace)
		}
	}()
	f()
}

func f() {
	var i *int
	*i = 0
}
