package main

import "fmt"

// Refactor of the 1_pipeline program to pass channels in functions
// Channels are passed by reference. Caller and callee refer to the same data structure.

func main() {

	// Unbuffered channel declaration
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go squarer(naturals, squares)

	printer(squares)
}

/*
	Unidirectional channels are channels which their type indicates
	whether they are meant to only receive or only send values to the channel
*/

// counter receives a channel which implicitly converts the channel type to a 'send-only' type.
// Performing a receive operation results to a compile-time error.
func counter(n chan<- int) {
	for i := 1; i <= 10; i++ {
		n <- i
	}
	close(n)
}

// squarer receives two channels which implicitly converts to
// 'send-only' and 'receive-only' channels.
func squarer(n <-chan int, sq chan<- int) {
	for nm := range n {
		sq <- nm * nm
	}
	// close(n) -> Performing a close on 'receive-only' channel is a compile-time error.
	close(sq)
}

// printer receives a channel which implicitly converts to a 'receive-only' channel type.
// Performing a send operation results to a compile-time error.
func printer(sq <-chan int) {
	for i := range sq {
		fmt.Println(i)
	}
}
