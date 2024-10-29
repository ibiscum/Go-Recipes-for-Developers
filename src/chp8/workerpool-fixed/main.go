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

const poolSize = 50

func doWork(input Input) Output {
	// Do some work
	time.Sleep(100 * time.Millisecond)
	return Output{
		Value: fmt.Sprint(input.Value),
	}
}

var inputCnt = 0

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

func workerPoolWithConcurrentReader() {
	inputCnt = 0
	// Send inputs to the pool via inputCh
	inputCh := make(chan Input)
	// Receive outputs from the pool via outputCh
	outputCh := make(chan Output)

	// Create the pool of workers
	wg := sync.WaitGroup{}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for work := range inputCh {
				outputCh <- doWork(work)
			}
		}()
	}
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

	// This goroutine waits until all worker pool goroutines are done, then
	// closes the output channel
	go func() {
		// Wait until processing is complete
		wg.Wait()
		// Close the output channel so the reader goroutine can terminate
		close(outputCh)
	}()

	// This loop sends the inputs to the worker pool
	for {
		nextInput, done := getNextInput()
		if done {
			break
		}
		inputCh <- nextInput
	}
	// Close the input channel, so worker pool goroutines terminate
	close(inputCh)
	// Wait until the output channel is closed
	readerWg.Wait()
	// If we are here, all goroutines are done
}

func workerPoolWithConcurrentWriter() {
	inputCnt = 0
	// Send inputs to the pool via inputCh
	inputCh := make(chan Input)
	// Receive outputs from the pool via outputCh
	outputCh := make(chan Output)

	// Writer goroutine submits work to the worker pool
	go func() {
		for {
			nextInput, done := getNextInput()
			if done {
				break
			}
			inputCh <- nextInput
		}
		// Close the input channel, so worker pool goroutines terminate
		close(inputCh)
	}()

	// Create the pool of workers
	wg := sync.WaitGroup{}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for work := range inputCh {
				outputCh <- doWork(work)
			}
		}()
	}

	// This goroutine waits until all worker pool goroutines are done, then
	// closes the output channel
	go func() {
		// Wait until processing is complete
		wg.Wait()
		// Close the output channel so the reader goroutine can terminate
		close(outputCh)
	}()

	// Read results until outputCh is closed
	for result := range outputCh {
		// process result
		fmt.Println(result)
	}

}

func main() {
	workerPoolWithConcurrentReader()
	workerPoolWithConcurrentWriter()
}
