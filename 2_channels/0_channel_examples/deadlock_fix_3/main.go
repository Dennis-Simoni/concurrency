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
