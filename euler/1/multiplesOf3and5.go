package main

import "fmt"

// Find the sum of all the multiples of 3 or 5 below 1000.
// https://projecteuler.net/problem=1
func main() {
	multiples := findMultiples(1000)
	//printSlice(multiples)
	sum := sumMultiples(multiples)
	fmt.Println(sum)
}

func findMultiples(number int) []int {
	var multiples []int
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
		}
	}
	return multiples
}

func sumMultiples(multiples []int) int {
	var sum int
	for _, multiple := range multiples {
		sum = sum + multiple
	}

	return sum
}

func printSlice(mySlice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
}
