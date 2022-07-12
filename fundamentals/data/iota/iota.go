package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)
const (
	monday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	ram1 = 1 << iota
	ram2
	ram3
	ram4
	ram5
	ram6
	ram7
)

func main() {

	// Iota automatically increments by 1 when used.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(a)

	// Resets with a new declaration of consts.
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)

	fmt.Printf("%v is th")

	fmt.Println(ram7)
}
