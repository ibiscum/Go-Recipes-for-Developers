package main

import (
	"errors"
	"fmt"
)

func process(doPanic bool) (err error) {
	defer func() {
		r := recover()
		if e, ok := r.(error); ok {
			err = e
		}
	}()
	if doPanic {
		panic(errors.New("panic!"))
	}
	return nil
}

func main() {
	fmt.Printf("return value of process without panic: %v\n", process(false))
	fmt.Printf("return value of process with panic: %v\n", process(true))
}
