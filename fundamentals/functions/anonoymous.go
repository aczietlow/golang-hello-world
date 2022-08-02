package main

import "fmt"

func main() {
	far()

	// Define an anonymous function, and call it by appending with "()" like with declared functions
	// Throwing defer in for funze
	defer func() {
		fmt.Println("using an anonymous function.")
	}()

	func(x int) {
		fmt.Printf("The meaning of life in the universe is %v\n", x)
	}(42)

}

func far() {
	fmt.Println("hello")
}
