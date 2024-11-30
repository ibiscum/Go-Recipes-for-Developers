package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func cancelableFunc(ctx context.Context) error {
	for {
		if err := ctx.Err(); err != nil {
			fmt.Println("cancelableFunc canceled")
			return err
		}
		fmt.Println("cancelableFunc sleeping")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func cancelableGoroutine1(ctx context.Context) {
	for {
		if err := ctx.Err(); err != nil {
			fmt.Println("cancelableGoroutine1 canceled")
			return
		}
		fmt.Println("cancelableGoroutine1 sleeping")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func cancelableGoroutine2(ctx context.Context) {
	for {
		if err := ctx.Err(); err != nil {
			fmt.Println("cancelableGoroutine2 canceled")
			return
		}
		fmt.Println("cancelableGoroutine2 sleeping")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	cancelable, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Canceling...")
		cancel()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cancelableGoroutine1(cancelable)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cancelableGoroutine2(cancelable)
	}()
	fmt.Println(cancelableFunc(cancelable))
	wg.Wait()
}
