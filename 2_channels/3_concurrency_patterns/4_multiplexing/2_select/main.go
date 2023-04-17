package main

import (
	"fmt"
	"os"
	"time"
)

// Another rocket launch counter implemented with the use of time.After

func main() {

	abort := make(chan struct{})

	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown")

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("rocket launched")
	case <-abort:
		fmt.Println("launch aborted")
		return
	}
}
