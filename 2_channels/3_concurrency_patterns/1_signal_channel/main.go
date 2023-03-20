package main

import (
	"fmt"
	"time"
)

// This program is using a channel as a signal mechanism.
func main() {
	done := make(chan struct{})
	fmt.Println("Program start time:", time.Now().Format("15:01:05"))

	// goroutine running on background
	go func() {
		// mock some slow operation
		time.Sleep(5 * time.Second)
		// signal job is done
		done <- struct{}{}
	}()

	// main goroutine blocks until a send operation is performed on the channel
	<-done

	fmt.Println("Program end time:", time.Now().Format("15:01:05"))
	// main terminates
}
