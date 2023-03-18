package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	Recap:
	We saw that unbuffered channels block as soon as they have performed their
	operation (e.g send) until the corresponding operation is performed (e.g receive)
	on the same channel.

	Buffered channels:
      Buffered channels introduce the idea of a queue.
	  On the declaration of a buffered channel we can determine the size of the queue.

	- The big difference of buffered channels is in the blocking mechanism.
		- Send operation:
		As long as the queue has available space, send operations will add to the queue and not block the goroutine.
		Once the maximum space has been occupied by values then the goroutine blocks until a receive operation is
      	performed on the channel
		- Receive operation:
		Retrieves values from the front of the queue, as long as there are values present in the queue.
		The blocking occurs when a receive is performed on an empty channel.
*/

// The following program sends a http GET request to three different urls
// and prints the fastest host.

func main() {

	// declaring a buffered channel of size 3
	// that means we can perform upto 3 sends without the goroutine being blocked
	ch := make(chan string, 3)

	// first send
	go func() {
		ch <- get("https://www.twitter.com")
	}()

	// second send
	go func() {
		ch <- get("https://www.amazon.com")
	}()

	// third send
	go func() {
		ch <- get("https://www.google.com")
	}()

	// another send at this point would have blocked the goroutine performing it.

	// Now, we perform a receive once, which will return the first value added to the channel
	fmt.Println("fastest response came from:", <-ch)

	// had we used an unbuffered channel in this case, the slower goroutines would have been stuck
	// trying to send their responses on a channel from which no goroutine will ever receive, this
	// situation is called 'goroutine leak' and it would be a bug

	// Leaked goroutines are not automatically collected, so it is important to make sure that
	// goroutines terminate themselves when no longer needed.
}

func get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Request.Host
}
