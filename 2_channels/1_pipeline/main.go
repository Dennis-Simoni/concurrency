package main

import "fmt"

func main() {

	// main goroutine.

	// declare unbuffered channels
	naturals := make(chan int)
	squares := make(chan int)

	// new goroutine running in background
	// produce numbers and add them to the channel
	go func() {
		for i := 1; i <= 10; i++ {
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
		for n := range naturals {
			squares <- n * n
		}

		// when this goroutine is done adding data to the channel
		// it can indicate that no more data will be provided by
		// closing the channel
		close(squares)
	}()

	// main goroutine

	// drain the squares channel and terminate
	for i := range squares {
		fmt.Println(i)
	}
}
