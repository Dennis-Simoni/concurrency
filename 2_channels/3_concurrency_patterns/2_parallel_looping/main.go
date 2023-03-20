package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// The following program demonstrates a pattern to loop in parallel
// assumptions:
// the order of work execution is irrelevant
// the size of the work to process is unknown

func main() {

	n := make(chan int)

	// random size
	size := rand.Intn(5000-1000) + 1000

	// in a new goroutine we populate n with integers
	go func() {
		for i := 1; i <= size; i++ {
			n <- i
		}
		// close the channel when done
		close(n)
	}()

	// pass the channel so the work function can do some work with the values in the channel
	work(n)
}

// receive a channel of integers that need to be squared
func work(ch <-chan int) {

	// the wait group acts as a counter to help keep track of how many goroutines we have running
	var wg sync.WaitGroup

	squared := make(chan int)

	// the channel size is unknown right now, we are just ranging until there is no more work to do
	for i := range ch {

		wg.Add(1) // increment the counter by one

		go func(i int) {
			//when the work is done decrement the counter by 1
			defer wg.Done()

			// mock some work
			time.Sleep(2 * time.Second)
			squared <- i * i
		}(i)
	}

	// closer goroutine
	go func() {
		// Wait blocks until the WaitGroup counter is zero.
		wg.Wait()

		// once we have reached zero, meaning no more goroutines are active
		// we are closing the squared channel to indicate no more values
		// will be added in
		close(squared)
	}()

	// drain squared until closed.
	for s := range squared {
		fmt.Println(s)
	}
}
