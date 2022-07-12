package main

import (
	"fmt"
)

func main() {
	// Anonymous struct.
	// Anonymous structs do not have a name.
	// If the struct isn't ever be reused, it doesn't need a named declaration.
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   42,
	}
	fmt.Println(p1)
}
