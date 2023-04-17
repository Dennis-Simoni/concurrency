package main

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch5/links"
	"log"
	"os"
)

// We can limit parallelism using buffered channels to model
// a concurrency primitive called 'counting semaphore'

var tokens = make(chan struct{}, 20)

func main() {

	worklist := make(chan []string)

	// counter to keep track the number of goroutines
	// so that we know when to terminate the program
	var n int

	n++

	// new goroutine adds to the channel
	// values passed in as program arguments
	go func() {
		worklist <- os.Args[1:]
	}()

	// seen is a set that helps avoid
	// crawling on a url more than once.
	seen := make(map[string]bool)

	for ; n > 0; n-- { // break the loop once all goroutines are finished
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true

				n++

				// new goroutine adds to the worklist
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	// if this operation causes the channel buffer to reach
	// its limit (20), then it will block until space is made up.
	list, err := links.Extract(url)
	<-tokens // release the token when done extracting
	if err != nil {
		log.Println(err)
	}

	return list
}
