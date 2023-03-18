package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("Process number %d is running", i)
		}
		// when we know we are done sending values to the channel then we can indicate that
		// by using the built-in close() function
		close(ch)
	}()

	// Recall: The loop for i := range ch receives values from the channel repeatedly until it is closed.
	for i := range ch {
		fmt.Println(i)
	}

	// main returns
}
