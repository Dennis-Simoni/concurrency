package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// This program is an echo server which tries to simulate the reverberations of a real echo.

/*
	Run instructions:
	- cd to this directory
	- go build main.go
	- ./main
	- open a terminal tab, use a tool like netcat to manipulate the network connection
	- nc localhost 8080
		- type in a word for the server to echo, ie golang
		- type another word, ie that is cool
*/

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		// for every new connection we want a new goroutine so to handle
		// more than one client at a time.
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	// For every client input, we want to execute echo concurrently
	// rather than having to wait for each word to be echoed back
	// so to achieve a more realistic echo
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	err := c.Close()
	if err != nil {
		return
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	_, err := fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	if err != nil {
		return
	}
	time.Sleep(delay)
	_, err = fmt.Fprintln(c, "\t", shout)
	if err != nil {
		return
	}
	time.Sleep(delay)
	_, err = fmt.Fprintln(c, "\t", strings.ToLower(shout))
	if err != nil {
		return
	}
}
