package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	// convert the int to type ByteCounter
	*c += ByteCounter(len(p))

	// Don't actually do anything with the bytes, just return the length of the slice, or number of bytes passed.
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c = WordCounter(0)
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

// Adding a receiver to type T that accepts *T means this method can only be called on a variable, and not an
// unassigned type or empty interface{}
func (c *LineCounter) Write(p []byte) (int, error) {
	*c = LineCounter(0)
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	// Interfaces create a contract between a concrete type and a caller/
	// In go, interfaces are satisfied implicitly. A concrete type that satisfies all necessary methods will suffice.

	var c ByteCounter
	c.Write([]byte("Hello World"))
	fmt.Println(c) // 11

	// fmt.Fprintf() File print f can write to a file, but all the w HAS to do is satisfy io.Writer{} which implements:
	// Write(p []byte) (int, error)
	// In this case could pass a file pointer, or use the type we defined above that implements Write() and satisfies this interface
	c = 0
	var name = "Floyd"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c) // 12

	story := "The is the song that never ends.\n Yes it goes on and on my friends.\n Some people started singing it not knowing what it was.\n"
	// Write a concrete type that counts the words when the Write method is called. Again disregard the bytes.
	var words WordCounter
	fmt.Fprintf(&words, story)
	fmt.Println(words)

	var lines LineCounter
	fmt.Fprintf(&lines, story)
	fmt.Println(lines)

}
