package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func longRunningGoroutine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Process some data
		// Check context cancelation
		select {
		case <-ctx.Done():
			// Context canceled
			fmt.Println("Canceled")
			return
		default:
		}
		// Continue computation
		fmt.Println("Computing...")
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}

func main() {
	ctx := context.Background()
	timeoutable, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go longRunningGoroutine(timeoutable, &wg)
	wg.Add(1)
	go longRunningGoroutine(timeoutable, &wg)
	wg.Wait()
}
