package main

import (
	"fmt"
	"os"
	"time"
)

// This program is a rocket launch counter.

func main() {

	// abort acts as a signal channel to abort the process
	abort := make(chan struct{})

	// a user input will cause to abort the rocket launch
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// The time.Tick function returns a channel on which it sends
	// events periodically
	fmt.Println("commencing countdown")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {

		// the general form of a select statement shown below, like a switch statement
		// it has a number of cases and an optional default.

		// a select waits until communication for some case is ready to proceed.
		// it then performs the communication and the case's associated statements
		// the other communications do not happen.

		// a select {} with no cases waits forever.

		select {
		case <-abort:
			fmt.Println("launch aborted")
			return
		case <-tick:
			fmt.Println(countdown)
		}
	}
}
