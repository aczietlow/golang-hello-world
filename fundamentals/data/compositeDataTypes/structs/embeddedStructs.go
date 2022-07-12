package main

import "fmt"

// Structs are just a collection of fields.
type humanPerson struct {
	first string
	last  string
}

type developer struct {
	humanPerson
	cupsOfCoffee int
}

func main() {
	dev := developer{
		humanPerson: humanPerson{
			first: "Larry",
			last:  "Ellison",
		},
		cupsOfCoffee: 4,
	}

	person2 := humanPerson{
		first: "Jill",
		last:  "Valentine",
	}

	fmt.Println(dev)
	fmt.Println(person2)

	fmt.Println(dev.cupsOfCoffee)
	// Can access inner type (person in this case) from the outer type.R
	fmt.Println(dev.first)
	// Optional to be more explicit and namespace with the inner type.
	// Subfields are referred to as promoted fields. They're promoted to the outer type.
	fmt.Println(dev.humanPerson.last)
}
