package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// This is a sequential server in which can only handle one client at a time
// The server responds to the client with a timestamp every second.

/*
	 	Run instructions:
			- cd to this directory
			- go build main.go
			- ./main
			- in a new terminal tab use a tool like netcat to manipulate the network connection
			- nc localhost 8080
*/

// if we try to open another netcat client in a new terminal we will have to wait until the
// current client terminates its connection.

/*
	Run instructions:
		- in a new terminal run nc localhost 8080 (note it will be stale until the next step)
		- terminate currently running client with control + C
		- observe how the second client now has established the connection
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

		handleConn(conn)
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
