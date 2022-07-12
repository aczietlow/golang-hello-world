package main

import (
	"fmt"
)

func main() {
	// Create your own type person which will have an underlying type of "struct" so that it can store the following data:
	// first name
	// last name
	// favorite ice cream flavors
	// Create two VALUES of TYPE person. Print out the values ranging over the elements in the slice which stores the favorite flavors
	one()

	// @TODO comeback to here.
}

func one() {
	type person struct {
		first   string
		last    string
		flavors []string
	}
	p1 := person{
		first:   "Jim",
		last:    "Dandy",
		flavors: []string{"chocolate", "vanilla"},
	}
	fmt.Println(p1)
	p2 := person{
		first:   "alaina",
		last:    "Zietlow",
		flavors: []string{"vanilla", "mouse tracks", "almond milk"},
	}
	fmt.Println(p2.first)

	for _, flavor := range p2.flavors {
		fmt.Println(flavor)
	}
}
