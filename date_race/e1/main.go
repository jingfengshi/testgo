package main

import (
	"fmt"
	"sync"
)

var counter int

func main(){
	const grs =10
	var wg sync.WaitGroup
	wg.Add(grs)
	// Create two goroutines.
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {

				// Capture the value of Counter.
				value := counter

				// Increment our local value of Counter.
				value++

				// Store the value back into Counter.
				counter = value
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}
