package main

import "fmt"

func main() {
	// Ideally use the most limiting scope possible, per norm
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
}

func incrementor() func() int {
	var x int
	// This anonymous func has access to x, even though it's defined outside its scope.
	// "Encloses a variable"
	return func() int {
		x++
		return x

	}
}
