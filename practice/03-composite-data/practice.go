package main

import "fmt"

func main() {

	// Using a composite literal
	// - create an array which holds 5 values of type int
	// - assign values to each index position
	// Range over the array and print the values out
	// using format printing
	// - print out the type of array
	one()
	// Wouldn't normally create an array

	// Using a composite literal
	// - create a slice of type int
	// - assign 10 values
	// Range over the slice and print the values out
	// using format printing
	// - print out the type of slice
	two()

	// Take code from previous example
	// slice the slice from the previous example
	// display the following
	// [2, 4, 8, 16, 32]
	// [64, 128, 256, 512, 1028]
	// [8, 16, 32, 64, 128, 256]
	// [4, 8, 16, 32, 64, 128]
	three()
}

func one() {
	a := [5]int{2, 4, 8, 16, 32}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of a is: %T\n", a)
}

func two() {
	s := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1028}
	fmt.Println(len(s))
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)
}

func three() {
	s := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1028}
	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:8])
	fmt.Println(s[1:7])
}
