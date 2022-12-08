package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	readTheThing(os.Stdout, conn)
}

func readTheThing(destination io.Writer, source io.Reader) {
	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal(err)
	}
}
