package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// This is the concurrent version of the clock server.

/*
	 	Run instructions:
			- cd to this directory
			- go build main.go
			- ./main
			- open a couple of new terminal tabs, use a tool like netcat to manipulate the network connection
			- nc localhost 8080

Note how n clients now are able to receive a response from the server, something that was not possible
in our sequential version.
*/
func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ready to accept connections:")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}

		// execute a new goroutine for every new connection / client.
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
