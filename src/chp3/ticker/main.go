package main

import (
	"fmt"
	"time"
)

func everySecond(f func(), done chan struct{}) {
	// Create a new ticker with a 1 second period
	ticker := time.NewTicker(time.Second)
	start := time.Now()
	// Stop the ticker once we're done
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			// Call the function
			fmt.Println(time.Since(start).Milliseconds())
			f()
		}
	}
}

func main() {
	cnt := 0
	done := make(chan struct{})
	everySecond(func() {
		switch cnt {
		case 0: // First call lasts 10 msecs
			time.Sleep(10 * time.Millisecond)
			cnt++
		case 1: // Second call lasts 1.5seconds
			time.Sleep(1500 * time.Millisecond)
			cnt++
		default:
			// Remaining calls are quick, and stop after 5 calls
			cnt++
			if cnt == 5 {
				close(done)
			}
		}
	}, done)

}
