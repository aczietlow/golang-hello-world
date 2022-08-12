package main

import "fmt"

func main() {
	// Nothing special.
	fact := factorial(8)
	fmt.Println(fact)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
