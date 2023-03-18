package main

import (
	"fmt"
	"time"
)

// This program produces a deadlock to demonstrate the effects of
// goroutines and (unbuffered) channels can have if not synchronised/implemented properly

func main() {

	// channel declaration of type: chan string
	// channels can have other types such as struct, int, float...
	ch := make(chan string)

	// channels are used as a means of communication between goroutines.
	// in this example we are running a goroutine which adds values to a channel
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)

			// This is the syntax for a send operation to the channel
			// this goroutine blocks from progressing until a receive '<-ch' operation
			// is performed to the same channel
			ch <- fmt.Sprintf("Process number %d is running", i)
		}
	}()

	// The way channels work cause the goroutines to block.
	// in this particular example the main goroutine blocks
	// from progressing until there is a send operation performed on the channel
	for p := range ch {
		fmt.Println(p)
	}

	// Explanation of hidden bug:

	// The loop for i := range ch receives values from the channel repeatedly until it is closed.
	// The main goroutine is blocked forever in this example, the reason being the goroutine that produces values has reached its end
	// without closing the channel and as it stands, there will never be another send operation to unblock main from progressing
}

/*
	How unbuffered channels work:
	A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes the corresponding
	receive operation on the same channel. At which point the value is transmitted and both goroutines may continue.

	Conversely, if the receive operation was attempted first, the receiving goroutine is blocked until another goroutine
	performs a send on the same channel
*/
