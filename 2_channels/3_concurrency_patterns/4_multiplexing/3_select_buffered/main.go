package main

import "fmt"

// Example select statement on a buffered channel

func main() {

	// channel with buffer of 1
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {

		// as the channel is alternately empty only one case can proceed

		// in the first iteration receiving is not ready (chan is empty)
		// so, the send operation is chosen.

		// in the second iteration the send operation can't perform (chan is full)
		// so the receive operation is chosen

		select {
		case x := <-ch:
			// receiving from a buffered channel blocks
			// the goroutine only if the channel is empty.
			fmt.Print(x, " ") // 0 2 4 6 8
		case ch <- i:
			// sending to a buffered channel blocks
			// the goroutine only if the channel is full.
		}
	}
}

// if multiple select cases are ready then one is picked at random
// which ensures every channel has an equal chance of being picked.

// Test your understanding: if we convert the channel to unbuffered, what is the result and why?

// Answer:

// sending and receiving to the channel happens on the main goroutine
// unbuffered channels block the moment a send/receive is performed.
// in this program, if we were to convert the channel to unbuffered
// one of the select statements would be picked at random and would
// block the main goroutine, leaving no chance to the corresponding
// operation to be performed, which leads to deadlock.

// Test your understanding: what if we increased the size of the buffer?

// Answer:

// increasing the buffer size makes the output nondeterministic as
// the select statement chooses randomly the operations
