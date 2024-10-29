package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type InputPayload struct {
	Id int
	// Add payload data fields here
}

type Stage2Payload struct {
	Id  int
	Err error
	// Stage2 data fields here
}

type Stage3Payload struct {
	Id  int
	Err error
	// Stage3 data fields here
}

type OutputPayload struct {
	Id  int
	Err error
}

func processData(id int) error {
	// Process data
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	// There may be an error
	if rand.Intn(100) < 10 {
		return fmt.Errorf("Processing failure for id: %d", id)
	}
	return nil
}

type PipelineError struct {
	Stage   int
	Payload any
	Err     error
}

func (p PipelineError) Error() string {
	return fmt.Sprintf("Pipeline error at stage: %d. Payload: %v. Cause: %s", p.Stage, p.Payload, p.Err)
}

func Stage1(input <-chan InputPayload) <-chan Stage2Payload {
	output := make(chan Stage2Payload)
	go func() {
		defer close(output)
		// Process all inputs
		for in := range input {
			// Process data
			err := processData(in.Id)
			if err != nil {
				err = PipelineError{
					Stage:   1,
					Payload: in,
					Err:     err,
				}
			}
			output <- Stage2Payload{
				Id:  in.Id,
				Err: err,
			}
		}
	}()
	return output
}

func Stage2(input <-chan Stage2Payload) <-chan Stage3Payload {
	output := make(chan Stage3Payload)
	go func() {
		defer close(output)
		// Process all inputs
		for in := range input {
			if in.Err != nil {
				output <- Stage3Payload{
					Id:  in.Id,
					Err: in.Err,
				}
				continue
			}
			// Process data
			err := processData(in.Id)
			if err != nil {
				err = PipelineError{
					Stage:   2,
					Payload: in,
					Err:     err,
				}
			}
			output <- Stage3Payload{
				Id:  in.Id,
				Err: err,
			}
		}
	}()
	return output
}

func Stage3(input <-chan Stage3Payload) <-chan OutputPayload {
	output := make(chan OutputPayload)
	go func() {
		defer close(output)
		// Process all inputs
		for in := range input {
			// Process data
			if in.Err != nil {
				output <- OutputPayload{
					Id:  in.Id,
					Err: in.Err,
				}
				continue
			}
			err := processData(in.Id)
			if err != nil {
				err = PipelineError{
					Stage:   3,
					Payload: in,
					Err:     err,
				}
			}
			output <- OutputPayload{
				Id:  in.Id,
				Err: err,
			}
		}
	}()
	return output
}

type SortQueue []OutputPayload

func (q SortQueue) Len() int { return len(q) }
func (q SortQueue) Less(i, j int) bool {
	return q[i].Id < q[j].Id
}

func (q SortQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *SortQueue) Push(x any) {
	*q = append(*q, x.(OutputPayload))
}

func (q *SortQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}

func order(input <-chan OutputPayload, bufsize int) <-chan OutputPayload {
	result := make(chan OutputPayload)
	queue := SortQueue{}
	go func() {
		defer close(result)
		for data := range input {
			// Add new item to the queue
			heap.Push(&queue, data)
			// If the queue grew enough, pop
			for len(queue) >= bufsize {
				result <- heap.Pop(&queue).(OutputPayload)
			}
		}
		for len(queue) > 0 {
			result <- heap.Pop(&queue).(OutputPayload)
		}
	}()

	return result
}

func fanIn(inputs []<-chan OutputPayload) <-chan OutputPayload {

	result := make(chan OutputPayload)

	// Listen to input channels in separate goroutines
	inputWg := sync.WaitGroup{}
	for inputIndex := range inputs {
		inputWg.Add(1)
		go func(index int) {
			defer inputWg.Done()
			for data := range inputs[index] {
				// Send the data to the output
				result <- data
			}
		}(inputIndex)
	}

	// When all input channels are closed, close the fan in ch
	go func() {
		inputWg.Wait()
		close(result)
	}()

	return result
}

func main() {
	inputCh := make(chan InputPayload)

	poolSize := 5
	outputs := make([]<-chan OutputPayload, 0)
	// All Stage1 goroutines listen to a single input channel
	for i := 0; i < poolSize; i++ {
		outputCh1 := Stage1(inputCh)
		outputCh2 := Stage2(outputCh1)
		outputCh3 := Stage3(outputCh2)
		outputs = append(outputs, outputCh3)
	}

	outputCh := order(fanIn(outputs), 15)

	// Feed input asynchronously
	go func() {
		defer close(inputCh)
		for i := 0; i < 1000; i++ {
			inputCh <- InputPayload{
				Id: i,
			}
		}
	}()

	// Read outputs
	for out := range outputCh {
		fmt.Println(out)
	}

}
