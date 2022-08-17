package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	goroutines = 100
)

func init(){
	rand.Seed(time.Now().UnixNano())
}


func main(){
	values :=make(chan int)

	var wg sync.WaitGroup

	wg.Add(goroutines)

	for gr :=0;gr < goroutines;gr++{
		go func() {
			defer wg.Done()

			n:=rand.Intn(100)

			if n%2 ==0{
				return
			}
			values <-n

		}()
	}

	go func(){
		wg.Wait()
		close(values)
	}()

	var nums []int
	for n:=range values{
		nums = append(nums,n)
	}

	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}