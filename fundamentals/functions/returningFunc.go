package main

import "fmt"

func main() {
	s1 := foo1()
	fmt.Println(s1)

	i1 := bar1()
	fmt.Printf("%T\n", i1)   // func() int
	fmt.Printf("%T\n", i1()) // 404
	// Cleaner way of writing this.
	fmt.Println(bar1()())
}

func foo1() string {
	s := "hello world"
	return s
}

// Returns a function that itself returns an int
func bar1() func() int {
	return func() int {
		return 404
	}

}
