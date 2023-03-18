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
	}()

	// one way of fixing the deadlock is not to range over the channel but range over the
	// number of sends we know we will receive

	// There is a better pattern to use however, see deadlock_fix_2

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

}
