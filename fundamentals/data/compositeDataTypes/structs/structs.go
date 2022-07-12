package main

import "fmt"

// https://go.dev/ref/spec#Struct_types
// Structs are just a collection of fields.
type person struct {
	first string
	last  string
}

func main() {
	// Create a value of different *types*
	person1 := person{
		first: "James",
		last:  "Bond",
	}

	// Not an instance. Creating a data STRUCTure of different types.
	person2 := person{
		first: "Chris",
		last:  "Zietlow",
	}

	fmt.Println(person1)
	fmt.Println(person2)

	// Access field with dot notation.
	fmt.Println(person1.first)
	fmt.Println(person2.last)
}
