package main

import (
	"fmt"
	"math"
)

func main() {
	variables()
}

func variables() {
	var day int // Declares int variable.
	fmt.Println("Been working on go for", day, "days.")
	day = 5
	fmt.Println("Been working on go for", day, "days.")

	// Declare multiple variables.
	var (
		name     string = "Chris"
		age      int    = 20
		location string = "BLT"
	)
	fmt.Println("Hola, me llama es", name, ". Soy", age, "anos y vivo en", location)

	// Declare multiple variables of the same type
	var width, height int
	fmt.Println("width is", width, "height is", height)
	width = 100
	height = 100
	fmt.Println("width is", width, "height is", height)

	// Short hand declaration.
	first_name, last_name := "Chris", "Zietlow"
	fmt.Println(first_name, last_name)

	// Variable types are inferred if not stated.
	a, b := 145.0, 152.8
	c := math.Min(a, b)
	fmt.Println(c)

}
