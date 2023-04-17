package main

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch5/links"
	"log"
	"os"
)

// This is an alternative pattern to limit concurrency

func main() {

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// new goroutine adds to the channel
	// values passed in as program arguments
	go func() {
		worklist <- os.Args[1:]
	}()

	// spin up 20 goroutines which will crawl on
	// deduplicated links
	for i := 0; i < 20; i++ {

		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)

				//
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true

				// push the link to the channel so
				// it can be crawled
				unseenLinks <- link
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
