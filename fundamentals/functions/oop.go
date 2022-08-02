package main

import "fmt"

// Interfaces and Polymorphism

type people struct {
	first string
	last  string
}

type agent struct {
	people
	ltk bool
}

// any struct that attaches the speak() method, will also be of type human
type human interface {
	speak() string
}

// You can use an empty interface for everything to get the interface.
// Everything value is of type empty interface (interface{}
type matter interface{}

type hotdog int

func check(h human) {
	// Struct.(type) will evaluate the type of the struct
	// This is an "assert". e.g. asset this variable is of a given type
	// only usable within switch statements
	switch h.(type) {
	// useful for accessing concrete type that is using this interface
	case people:
		fmt.Println("I am a human being, I was passed into check", h.(people).first)
	case agent:
		fmt.Println("I am a human being, I was passed into check", h.(agent).first)
	}
	fmt.Println("I am a human being, I was passed into check", h)
}

func (s agent) speak() string {
	return fmt.Sprintf("I am %v %v a secret agent.\n", s.first, s.last)
}

func (p people) speak() string {
	return fmt.Sprintf("I am %v %v a normal person \n", p.first, p.last)
}

func main() {
	// composite literal
	sa2 := agent{
		people: people{
			"james",
			"bond",
		},
		ltk: true,
	}

	fmt.Println(sa2.speak())

	sa1 := agent{
		people: people{
			"agent",
			"q",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Printf("type:%T\n", sa1)

	p1 := people{
		"Brandon",
		"Erlich",
	}

	// people is a human
	check(p1)
	// secret agent is also a human (because the agent struct implements the speak() method
	check(sa1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
