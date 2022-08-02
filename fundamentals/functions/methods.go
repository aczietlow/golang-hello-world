package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() string {
	return fmt.Sprintf("I am %v %v", s.first, s.last)
}

// func (r receiver) identifier (parameters) (return(s)) { ... }

func main() {
	// composite literal
	p1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}

	sa1 := secretAgent{
		person: person{
			"agent",
			"q",
		},
		ltk: true,
	}

	fmt.Println(sa1)

	fmt.Println(p1.speak())
}
