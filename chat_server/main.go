package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// The chat server lets several users broadcast textual messages
// to each other. There are four goroutine types in this program
// main, broadcaster, handleConn and clientWriter

// client is the outgoing message channel
type client chan<- string

var (
	entering = make(chan client) // channel of type chan<- string
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func main() {

	// main goroutine accepts new connections.
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {

	// clients records the current set of connected clients
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			// broadcast message to all clients' outgoing message channel
			for c := range clients {
				c <- msg
			}
		case c := <-entering:
			// update the client list with new client
			clients[c] = true
		case c := <-leaving:
			// update the client list by removing leaving client
			// close the client's channel
			delete(clients, c)
			close(c)
		}
	}
}

func handleConn(c net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(c, ch)

	cid := c.RemoteAddr().String()
	ch <- fmt.Sprintf("client identity: %s", cid)
	messages <- cid + ": has entered."
	entering <- ch

	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- cid + ": " + input.Text()
	}

	leaving <- ch
	messages <- cid + "has left."
	c.Close()
}

func clientWriter(c net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}
