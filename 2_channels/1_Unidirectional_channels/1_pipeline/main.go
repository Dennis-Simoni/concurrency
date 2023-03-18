package main

import "fmt"

// This is a pipeline program that uses channels to orchestrate the processing of some arbitrary work

func main() {

	// main goroutine.

	// declare unbuffered channels
	naturals := make(chan int)
	squares := make(chan int)

	// new goroutine running in background
	// produce numbers and add them to the channel
	go func() {
		for i := 1; i <= 10; i++ {

			// the goroutine blocks here until a receive operation is performed
			// to the same channel
			naturals <- i
		}

		// when this goroutine is done adding data to the channel
		// it indicates that no more data will be provided by
		// closing the channel
		close(naturals)
	}()

	// new goroutine running in background
	// drain the naturals channel and add
	// result of calculation to squares channel
	go func() {

		// the range is causing a receive operation, so the goroutine
		// which sends to the naturals channel can unblock
		for n := range naturals {

			// the goroutine blocks here until a recieve is performed
			// to the same channel
			squares <- n * n
		}

		// when this goroutine is done adding data to the channel
		// it can indicate that no more data will be provided by
		// closing the channel
		close(squares)
	}()

	// main goroutine

	// the range is causing a receive operation so the goroutine
	// which sends to the squares channel can unblock
	for i := range squares {

		// the main goroutine blocks here until a send is performed
		// to the squares channel
		fmt.Println(i)
	}

	// the main goroutine ends as the squares channel is closed.
	// recall range on channel happens until it is closed.
}
