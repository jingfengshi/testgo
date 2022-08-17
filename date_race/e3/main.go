package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

// mutex is used to define a critical section of code.
var mutex sync.Mutex

func main() {

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {

				// Only allow one goroutine through this critical section at a time.
				mutex.Lock()
				{
					// Capture the value of counter.
					value := counter

					// Increment our local value of counter.
					value++

					// Store the value back into counter.
					counter = value
				}
				mutex.Unlock()
				// Release the lock and allow any waiting goroutine through.
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}