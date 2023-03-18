package main

import (
	"fmt"
	"time"
)

// The spinner program demonstrates how two autonomous activities
// namely spinner and slowProcess can progress concurrently.

func main() {

	// main goroutine

	go spinner() // -> new goroutine running on the background

	// function running as part of the main goroutine
	// the program will execute until the slow process
	// is complete in which case the main goroutine terminates.
	// When this happens all goroutines are abruptly terminated
	// and the program exits.
	slowProcess()
}

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func slowProcess() {
	time.Sleep(10 * time.Second)
	fmt.Println("\nslow process complete")
}
