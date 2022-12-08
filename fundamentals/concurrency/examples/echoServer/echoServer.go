package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	Listener, err := net.Listen("tcp", "localhost:9002")
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
func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay * 5)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handleConnection(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), time.Duration(1*time.Second))
	}
	conn.Close()
}
