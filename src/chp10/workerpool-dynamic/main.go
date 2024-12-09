package main

import (
	"fmt"
	"sync"
	"time"
)

type Input struct {
	Value int
}

type Output struct {
	Value string
}

func doWork(input Input) Output {
	// Do some work
	time.Sleep(100 * time.Millisecond)
	return Output{
		Value: fmt.Sprint(input.Value),
	}
}

var inputCnt = 0

const maxPoolSize = 50

const maxInput = 5000

func getNextInput() (Input, bool) {
	if inputCnt >= maxInput {
		return Input{}, true
	}
	inputCnt++
	return Input{
		Value: inputCnt,
	}, false
}

func main() {
	inputCnt = 0
	// Receive outputs from the pool via outputCh
	outputCh := make(chan Output)
	// A semaphore to limit the pool size
	sem := make(chan struct{}, maxPoolSize)

	// Reader goroutine reads results until outputCh is closed
	readerWg := sync.WaitGroup{}
	readerWg.Add(1)
	go func() {
		defer readerWg.Done()
		for result := range outputCh {
			// process result
			fmt.Println(result)
		}
	}()

	// Create the workers as needed, but the number of active workers
	// are limited by the capacity of sem
	wg := sync.WaitGroup{}
	// This loop sends the inputs to workers, creating them as necessary
	for {
		nextInput, done := getNextInput()
		if done {
			break
		}
		wg.Add(1)
		// This will block if there are too many goroutines
		sem <- struct{}{}
		go func(inp Input) {
			defer wg.Done()
			defer func() {
				<-sem
			}()
			outputCh <- doWork(inp)
		}(nextInput)
	}

	// This goroutine waits until all worker pool goroutines are done, then
	// closes the output channel
	go func() {
		// Wait until processing is complete
		wg.Wait()
		// Close the output channel so the reader goroutine can terminate
		close(outputCh)
	}()

	// Wait until the output channel is closed
	readerWg.Wait()
	// If we are here, all goroutines are done
}
