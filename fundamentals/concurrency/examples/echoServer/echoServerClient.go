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
	defer conn.Close()
	go copyTheThing(os.Stdout, conn)
	copyTheThing(conn, os.Stdin)

}

func copyTheThing(destination io.Writer, source io.Reader) {
	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal(err)
	}
}
