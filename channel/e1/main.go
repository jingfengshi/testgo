package main

import (
	"fmt"
	"sync"
)

func main(){
	 share :=make(chan int)

	 var wg sync.WaitGroup
	 wg.Add(2)

	 go func() {
		 goroutine("Bill", share)
		 wg.Done()
	 }()

	go func() {
		goroutine("Joan", share)
		wg.Done()
	}()

	// Start the sharing.
	share <- 1

	// Wait for the program to finish.
	wg.Wait()
}


func goroutine(name string,share chan int){
	for{
		value,ok :=<-share
		if !ok{
			// If the channel was closed, return.
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Increment the value and send it
		// over the channel.
		share <- (value + 1)
	}
}