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
		close(ch)
	}()

	// Another way of fixing the deadlock is by using the comma ok idiom to validate there is something to be retrieved
	// from the channel before we perform a receive operation, the check confirms whether the channel is closed or not.

	for {
		if m, ok := <-ch; !ok {
			break
		} else {
			fmt.Println(m)
		}
	}

	// main returns
}

/*
	It is not necessary to close every channel when we are finished with it, it is only necessary to close a
	channel when it is important to tell the receiving goroutines that all data has been sent.

	A channel that the garbage collector determines to be unreachable will have its resources reclaimed whether
	closed or not.

	Subsequent calls to a closed channel yield the channel's zero value, ie chan int will return 0.
*/
