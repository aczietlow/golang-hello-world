package main

import "fmt"

const (
	a = iota
	b
	c
)
const (
	friday = iota
	saturday
	sunday
)

func main() {

	// Iota automatically increments by 1 when used.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Resets with a new declaration of consts.
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)
}
