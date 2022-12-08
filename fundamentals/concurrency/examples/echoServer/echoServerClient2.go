package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9002")
	if err != nil {
		log.Fatal(err)
	}
	// Create a channel of struct type.
	// The type isn't actually important here, as we're only using it for the blocking that comes from unbuffered channels
	done := make(chan struct{})

	defer conn.Close()
	go func() {
		io.Copy(os.Stdout, conn) // Ignore errors
		log.Println("done")
		done <- struct{}{}

	}()
	copyTheThing2(conn, os.Stdin)
	conn.Close()
	<-done
}

func copyTheThing2(destination io.Writer, source io.Reader) {
	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal(err)
	}
}
