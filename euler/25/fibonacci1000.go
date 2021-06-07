package main

import (
	"fmt"
	"math/big"
)

// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
// https://projecteuler.net/problem=25
func main() {
	fm, fn := big.NewInt(0), big.NewInt(1)
	digits, index := 0, 1

	for digits < 1000 {
		// Sums fm and fn, assigning to fm.
		fm.Add(fm, fn)
		digits = len(fm.String())
		index++
		fm, fn = fn, fm
	}

	fmt.Printf("The first occurence of %d digits is at the %d fib index\n", digits, index)
	println(fm.String())
}
