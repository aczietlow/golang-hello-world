package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	Listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := Listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(connection) // Do something with this.
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // ¯\_(ツ)_/¯ close the connection and move along.
		}
		time.Sleep(1 * time.Second)
	}
}
