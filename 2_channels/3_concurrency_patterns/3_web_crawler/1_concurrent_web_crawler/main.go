package main

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch5/links"
	"log"
	"os"
)

// This is a highly concurrent program which results to errors
// due to rate limiting.

func main() {
	worklist := make(chan []string)

	// new goroutine adds to the channel
	// values passed in as program arguments
	go func() {
		worklist <- os.Args[1:]
	}()

	// seen is a set that helps avoid
	// crawling on a url more than once.
	seen := make(map[string]bool)

	// main goroutine drains worklist
	// hence for as long as there are values
	// to retrieve, the main goroutine blocks here
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true

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
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}

	return list
}
