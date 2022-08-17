package main

import (
	"fmt"
	"math/rand"
	"time"
)

const  (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


func main(){
	values := make(chan int, goroutines)

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			values <- rand.Intn(1000)
		}()
	}

	// Set the value to the number of goroutines created.
	wait := goroutines

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var nums []int
	for wait > 0 {
		nums = append(nums, <-values)
		wait--
	}

	// Print the values in our slice.
	fmt.Println(nums)
}

